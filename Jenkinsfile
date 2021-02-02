pipeline {
    agent { docker { image 'golang' } }

    environment {
	THREADS=10
        TIMEOUT=3000
    }

    stages {
        stage('Build') {
            steps {
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                sh 'go test'
            }
        }
    }
}
