pipeline {
    agent any
    stages {
        stage('Pre Test'){
            echo 'Pulling Dependencies'

            sh 'go get -u github.com/flybikeGx/easy-timeout'
            sh 'cd $GOPATH/src/github.com/flybikeGx/easy-timeout'
        }
        stage('Test'){
            sh 'go test -v'
        }
    }
}
