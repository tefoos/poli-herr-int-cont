pipeline {
    agent any

    stages {
        stage('Clone') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/tefoos/poli-herr-int-cont'
            }
        }

        stage('Build') {
            steps {
                sh 'docker build -t poli-herr-int-cont-api .'
            }
        }

        stage('Test') {
            steps {
                sh 'docker run --rm -v /var/lib/docker/volumes/jenkins_home/_data/workspace/poli-herr-int-cont-pipeline:/app -w /app golang:1.24-alpine go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                sh 'docker compose up -d --build'
            }
        }
    }
}
