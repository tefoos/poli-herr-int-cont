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
                sh 'go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                sh 'docker compose up -d --build'
            }
        }
    }
}
