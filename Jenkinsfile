
pipeline {
    agent any

    stages {
        stage('SCM') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                go build
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                go test -v
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
            }
        }
    }
}
