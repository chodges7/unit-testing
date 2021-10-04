
pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test -v'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
            }
        }
    }
}
