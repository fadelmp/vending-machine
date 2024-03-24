pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'docker build -t perqara/vending-machine-image:latest .'
            }
        }
        stage('Push') {
            steps {
                sh 'docker push perqara/vending-machine-image:latest'
            }
        }
        stage('Deploy') {
            steps {
                sh 'kubectl apply -f deployment.yaml'
                sh 'kubectl apply -f service.yaml'
            }
        }
    }
}
