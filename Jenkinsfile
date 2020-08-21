pipeline {
    agent { docker { image 'golang' } }

    environment {
        XDG_CACHE_HOME = "/tmp/.cache"
    }

    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'printenv'
                sh 'go build hello.go'
            }
        }
        stage('Sanity check') {
            steps {
                input "Does the staging environment look ok?"
            }
        }
        stage('test') {
            steps {
                sh './hello'
            }
        }
    }
}

