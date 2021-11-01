
pipeline {
    agent any
    environment {
        GITHUB_TOKEN = credentials('GITHUB_TOKEN')
    }
    stages {
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go test'
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'goreleaser build --rm-dist'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Publish...'
                sh 'goreleaser release --rm-dist'
            }
        }
    }
}
