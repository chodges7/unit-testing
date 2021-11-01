
pipeline {
    agent any
    environment {
        GITHUB_TOKEN = credentials('github-token')
    }
    stages {
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'goreleaser build --rm-dist'
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
                echo 'Publish...'
                sh 'goreleaser release --rm-dist'
            }
        }
    }
}
