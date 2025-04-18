pipeline {
    agent any

    environment {
        APP_NAME = 'my-go-app'
        IMAGE_NAME = 'my-go-app:latest'
        CONTAINER_NAME = 'go-app'
        PORT = '8080'
    }

    stages {
        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $IMAGE_NAME .'
            }
        }

        stage('Run Container') {
            steps {
                sh '''
                docker rm -f $CONTAINER_NAME || true
                docker run -d --name $CONTAINER_NAME -p $PORT:8080 $IMAGE_NAME
                '''
            }
        }
    }
}
