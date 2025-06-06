pipeline {
  agent any
  environment {
    GOPROXY = 'https://goproxy.cn,direct'
  }
  tools {
    go 'go'
  }
  stages {
    stage('Clone') {
      steps {
        git(url: scm.userRemoteConfigs[0].url, branch: '$BRANCH_NAME', changelog: true, credentialsId: 'KK-github-key', poll: true)
      }
    }

    stage('Prepare') {
      steps {
        sh 'rm -rf ./output/*'
        sh 'make deps'
      }
    }

    stage('Linting') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make verify'
      }
    }

    stage('Compile') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          make verify-build
        '''.stripIndent())
      }
    }

    stage('Switch to current cluster') {
      when {
        anyOf {
          expression { BUILD_TARGET == 'true' }
          expression { DEPLOY_TARGET == 'true' }
        }
      }
      steps {
        sh 'cd /etc/kubeasz; ./ezctl checkout $TARGET_ENV'
      }
    }

    stage('Config target') {
      when {
        anyOf {
          expression { BUILD_TARGET == 'true' }
          expression { DEPLOY_TARGET == 'true' }
        }
      }
      steps {
        sh 'rm .apollo-base-config -rf'
        sh 'git clone https://github.com/NpoolPlatform/apollo-base-config.git .apollo-base-config'
        sh(returnStdout: false, script: '''
          PASSWORD=`kubectl get secret --namespace "kube-system" mysql-password-secret -o jsonpath="{.data.rootpassword}" | base64 --decode`
          kubectl exec --namespace kube-system mysql-0 -- mysql -h 127.0.0.1 -uroot -p$PASSWORD -P3306 -e "create database if not exists billing_manager;"

          username=`helm status rabbitmq --namespace kube-system | grep Username | awk -F ' : ' '{print $2}' | sed 's/"//g'`
          for vhost in `cat cmd/*/*.viper.yaml | grep hostname | awk '{print $2}' | sed 's/"//g' | sed 's/\\./-/g'`; do
            kubectl exec --namespace kube-system rabbitmq-0 -- rabbitmqctl add_vhost $vhost
            kubectl exec --namespace kube-system rabbitmq-0 -- rabbitmqctl set_permissions -p $vhost $username ".*" ".*" ".*"

            cd .apollo-base-config
            ./apollo-base-config.sh $APP_ID $TARGET_ENV $vhost
            ./apollo-item-config.sh $APP_ID $TARGET_ENV $vhost database_name billing_manager
            cd -
          done
        '''.stripIndent())
      }
    }

    stage('Unit Tests') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          devboxpod=`kubectl get pods -A | grep development-box | head -n1 | awk '{print $2}'`
          servicename="billing-middleware"

          kubectl exec --namespace kube-system $devboxpod -- make -C /tmp/$servicename after-test || true
          kubectl exec --namespace kube-system $devboxpod -- rm -rf /tmp/$servicename || true
          kubectl cp ./ kube-system/$devboxpod:/tmp/$servicename

          kubectl exec --namespace kube-system $devboxpod -- make -C /tmp/$servicename deps before-test test after-test
          kubectl exec --namespace kube-system $devboxpod -- rm -rf /tmp/$servicename

          swaggeruipod=`kubectl get pods -A | grep swagger | awk '{print $2}'`
          kubectl cp message/npool/*.swagger.json kube-system/$swaggeruipod:/usr/share/nginx/html || true
        '''.stripIndent())
      }
    }

    stage('Generate docker image for development') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make verify-build'
        sh 'DOCKER_REGISTRY=$DOCKER_REGISTRY make generate-docker-images'
      }
    }

    stage('Tag patch') {
      when {
        expression { TAG_PATCH == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e
          if [ 0 -eq $rc -a x"$revlist" != x ]; then
            tag=`git describe --tags $revlist`

            major=`echo $tag | awk -F '.' '{ print $1 }'`
            minor=`echo $tag | awk -F '.' '{ print $2 }'`
            patch=`echo $tag | awk -F '.' '{ print $3 }'`

            case $TAG_FOR in
              testing)
                patch=$(( $patch + $patch % 2 + 1 ))
                ;;
              production)
                patch=$(( $patch + 1 ))
                git reset --hard
                git checkout $tag
                ;;
            esac

            tag=$major.$minor.$patch
          else
            tag=0.1.1
          fi
          git tag -a $tag -m "Bump version to $tag"
        '''.stripIndent())

        withCredentials([gitUsernamePassword(credentialsId: 'KK-github-key', gitToolName: 'git-tool')]) {
          sh 'git push --tag'
        }
      }
    }

    stage('Tag minor') {
      when {
        expression { TAG_MINOR == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e
          if [ 0 -eq $rc -a x"$revlist" != x ]; then
            tag=`git describe --tags $revlist`

            major=`echo $tag | awk -F '.' '{ print $1 }'`
            minor=`echo $tag | awk -F '.' '{ print $2 }'`
            patch=`echo $tag | awk -F '.' '{ print $3 }'`

            minor=$(( $minor + 1 ))
            patch=1

            tag=$major.$minor.$patch
          else
            tag=0.1.1
          fi
          git tag -a $tag -m "Bump version to $tag"
        '''.stripIndent())

        withCredentials([gitUsernamePassword(credentialsId: 'KK-github-key', gitToolName: 'git-tool')]) {
          sh 'git push --tag'
        }
      }
    }

    stage('Tag major') {
      when {
        expression { TAG_MAJOR == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e
          if [ 0 -eq $rc -a x"$revlist" != x ]; then
            tag=`git describe --tags $revlist`

            major=`echo $tag | awk -F '.' '{ print $1 }'`
            minor=`echo $tag | awk -F '.' '{ print $2 }'`
            patch=`echo $tag | awk -F '.' '{ print $3 }'`

            major=$(( $major + 1 ))
            minor=0
            patch=1

            tag=$major.$minor.$patch
          else
            tag=0.1.1
          fi
          git tag -a $tag -m "Bump version to $tag"
        '''.stripIndent())

        withCredentials([gitUsernamePassword(credentialsId: 'KK-github-key', gitToolName: 'git-tool')]) {
          sh 'git push --tag'
        }
      }
    }

    stage('Generate docker image for testing or production') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e
          if [ 0 -eq $rc -a x"$revlist" != x ]; then
            tag=`git describe --tags $revlist`
            git reset --hard
            git checkout $tag
          fi
        '''.stripIndent())
        sh 'make verify-build'
        sh 'DEVELOPMENT=other DOCKER_REGISTRY=$DOCKER_REGISTRY make generate-docker-images'
      }
    }

    stage('Release docker image for development') {
      when {
        expression { RELEASE_TARGET == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          branch=latest
          if [ "x$BRANCH_NAME" != "xmaster" ]; then
            branch=`echo $BRANCH_NAME | awk -F '/' '{ print $2 }'`
          fi
          set +e
          docker images | grep billing-middleware | grep $branch
          rc=$?
          set -e
          if [ 0 -eq $rc ]; then
            DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker-images
          fi
          images=`docker images | grep entropypool | grep billing-middleware | grep none | awk '{ print $3 }'`
          for image in $images; do
            docker rmi $image -f
          done
        '''.stripIndent())
      }
    }

    stage('Release docker image for testing') {
      when {
        expression { RELEASE_TARGET == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e

          if [ 0 -eq $rc -a x"$revlist" != x ]; then
            tag=`git tag --sort=-v:refname | grep [1\\|3\\|5\\|7\\|9]$ | head -n1`
            set +e
            docker images | grep billing-middleware | grep $tag
            rc=$?
            set -e
            if [ 0 -eq $rc ]; then
              DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker-images
            fi
          fi
        '''.stripIndent())
      }
    }

    stage('Release docker image for production') {
      when {
        expression { RELEASE_TARGET == 'true' }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          taglist=`git rev-list --tags`
          rc=$?
          set -e

          if [ 0 -eq $rc -a x"$taglist" != x ]; then
            tag=`git tag --sort=-v:refname | grep [0\\|2\\|4\\|6\\|8]$ | head -n1`
            set +e
            docker images | grep billing-middleware | grep $tag
            rc=$?
            set -e
            if [ 0 -eq $rc ]; then
              DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker-images
            fi
          fi
        '''.stripIndent())
      }
    }

    stage('Deploy for development') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { TARGET_ENV ==~ /.*development.*/ }
      }
      steps {
        sh(returnStdout: false, script: '''
          branch=latest
          if [ "x$BRANCH_NAME" != "xmaster" ]; then
            branch=`echo $BRANCH_NAME | awk -F '/' '{ print $2 }'`
          fi
          sed -i "s/billing-middleware:latest/billing-middleware:$branch/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          sed -i "s/uhub.service.ucloud.cn/$DOCKER_REGISTRY/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          if [ "x$REPLICAS_COUNT" == "x" ];then
            REPLICAS_COUNT=2
          fi
          sed -i "s/replicas: 2/replicas: $REPLICAS_COUNT/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Deploy for testing') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { TARGET_ENV ==~ /.*testing.*/ }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e
          if [ ! 0 -eq $rc -o x"$revlist" == x]; then
            exit 0
          fi
          tag=`git tag --sort=-v:refname | grep [1\\|3\\|5\\|7\\|9]$ | head -n1`

          git reset --hard
          git checkout $tag
          sed -i "s/billing-middleware:latest/billing-middleware:$tag/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          sed -i "s/uhub.service.ucloud.cn/$DOCKER_REGISTRY/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          if [ "x$REPLICAS_COUNT" == "x" ];then
            REPLICAS_COUNT=2
          fi
          sed -i "s/replicas: 2/replicas: $REPLICAS_COUNT/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          sed -i "s/imagePullPolicy: Always/imagePullPolicy: IfNotPresent/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Deploy for production') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { TARGET_ENV ==~ /.*production.*/ }
      }
      steps {
        sh(returnStdout: false, script: '''
          set +e
          taglist=`git rev-list --tags`
          rc=$?
          set -e
          if [ ! 0 -eq $rc -o x"$revlist" == x]; then
            exit 0
          fi
          tag=`git tag --sort=-v:refname | grep [0\\|2\\|4\\|6\\|8]$ | head -n1`
          git reset --hard
          git checkout $tag
          sed -i "s/billing-middleware:latest/billing-middleware:$tag/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          sed -i "s/uhub.service.ucloud.cn/$DOCKER_REGISTRY/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          if [ "x$REPLICAS_COUNT" == "x" ];then
            REPLICAS_COUNT=2
          fi
          sed -i "s/replicas: 2/replicas: $REPLICAS_COUNT/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          sed -i "s/imagePullPolicy: Always/imagePullPolicy: IfNotPresent/g" cmd/billing-middleware/k8s/02-billing-middleware.yaml
          make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Post') {
      steps {
        // Assemble vet and lint info.
        // warnings parserConfigurations: [
        //   [pattern: 'govet.txt', parserName: 'Go Vet'],
        //   [pattern: 'golint.txt', parserName: 'Go Lint']
        // ]

        // sh 'go2xunit -fail -input gotest.txt -output gotest.xml'
        // junit "gotest.xml"
        sh 'echo Posting'
      }
    }
  }
  post('Report') {
    fixed {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh fixed')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/success_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
    success {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh successful')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/success_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
    failure {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh failure')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
    aborted {
      script {
        sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh aborted')
     }
      script {
        // env.ForEmailPlugin = env.WORKSPACE
        emailext attachmentsPattern: 'TestResults\\*.trx',
        body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
        mimeType: 'text/html',
        subject: currentBuild.currentResult + " : " + env.JOB_NAME,
        to: '$DEFAULT_RECIPIENTS'
      }
     }
  }
}
