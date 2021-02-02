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
	
       stage('Package') {
           steps {
               sh 'docker build -t rmauge/webby .'
           }
       }
       
       stage('Deploy') {
           steps {
               sh 'docker run -p 8888:8080 --name "webby" rmauge/webby'
           }
       }
    }
}
