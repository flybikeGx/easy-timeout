pipeline {
    agent any
    stages {
        stage('Pre Test'){
            steps {
                echo 'Pulling Dependencies'

                sh 'go get -u github.com/flybikeGx/easy-timeout'
                sh 'cd $GOPATH/src/github.com/flybikeGx/easy-timeout'
            }

        }
        stage('Test'){
            steps {
                sh 'go test -v'
            }
        }
    }
}
