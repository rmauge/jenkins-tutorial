pipeline {
    agent { docker { image 'golang' } }

    environment {
	THREAD=10
        TIMEOUT=3000
    }

    stages {
        stage('build') {
            steps {
                sh 'echo "Step 1" THREADS=${THREADS}'
                sh 'go version'
            }
        }
    }
}
