pipeline {
    agent any
    tools { go '1.20' }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    // echo 'Running linting'
                    // sh 'golangci-lint run'
                    // echo 'Running test'
                    // sh 'cd test && go test -v'
                }
            }
        }
        
    }
}