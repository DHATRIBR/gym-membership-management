pipeline {
    agent any

    environment {
        STAGING_SERVER = 'user@staging-server-ip'
        PROD_SERVER = 'user@prod-server-ip'
    }

    stages {
        stage('Checkout') {
            steps {
                echo '🔄 Cloning repository...'
                git branch: 'main', url: 'https://github.com/DHATRIBR/gym-membership-management.git'
            }
        }

        stage('Build') {
            steps {
                echo '🏗️ Building the application...'
                sh 'go mod tidy'  // Make sure all dependencies are fetched
                sh 'go build -o myapp'  // Build the Go application
            }
        }

        stage('Test') {
            steps {
                echo '🧪 Running unit tests...'
                sh 'go test -v ./...'
            }
        }

        stage('Deploy to Staging') {
            steps {
                echo '🚀 Deploying to staging environment...'
            }
        }

        stage('Approval') {
            steps {
                input message: 'Ready to deploy to Production?'
            }
        }

        stage('Deploy to Production') {
            steps {
                echo '🚀 Deploying to production environment...'
            }
        }
    }

    post {
        success {
            echo '✅ Pipeline completed successfully!'
        }
        failure {
            echo '❌ Pipeline failed.'
        }
        always {
            echo '📦 Clean-up actions if needed.'
        }
    }
}

