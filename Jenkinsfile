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
                echo '🚀 Deploying to staging...'
                sh '''
                    mkdir -p staging
                    cp myapp staging/
                    echo "Deployed to STAGING at $(date)" >> staging/deploy.log
                '''
            }
        }

        stage('Deploy to Production') {
            steps {
                input message: 'Promote to Production?', ok: 'Deploy'
                echo '🚀 Deploying to production...'
                sh '''
                    mkdir -p production
                    cp myapp production/
                    echo "Deployed to PRODUCTION at $(date)" >> production/deploy.log
                '''
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

