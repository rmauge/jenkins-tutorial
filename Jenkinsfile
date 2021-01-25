pipeline {
    agent { docker { image 'golang' } }

    environment {
	THREADS=10
        TIMEOUT=3000
    }

    stages {
        stage('build') {
            steps {
                sh 'echo "Step 1" THREADS=${THREADS}'
                sh 'go version'
            }
        }

        stage('Sanity check') {
            steps {
                input "Does the staging environment look ok?"
            }
        }
    }
}
