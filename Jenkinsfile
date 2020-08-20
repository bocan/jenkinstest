pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'printenv'
                go build hello.go
            }
        }
        stage('test') {
            steps {
                hello
            }
        }
    }
}

