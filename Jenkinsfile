pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'echo "Step 1"'
                sh 'go version'
            }
        }
    }
}
