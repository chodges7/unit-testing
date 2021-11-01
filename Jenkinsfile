
pipeline {
    agent any
    environment {
        GITHUB_TOKEN = credentials('GITHUB_TOKEN')
    }
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
                sh 'go test example.go example_test.go drone.go drone_test.go -v'
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
