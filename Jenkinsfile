pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'printenv'
                sh 'go build hello.go'
            }
        }
        stage('test') {
            steps {
                sh './hello'
            }
        }
    }
}

