def secrets = [
  [path: 'kv/brantas/main', engineVersion: 2, secretValues: []],
]
def configuration = [vaultUrl: 'https://vault.ubed.dev',  vaultCredentialId: 'vaultapprole', engineVersion: 2]
pipeline {
    agent any
    environment{
        DOCKERHUB_CREDS = credentials('Dockerhub')
    }
    stages {
         stage('Clone Repo') {
             when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
                checkout scm
                sh '''#!/bin/bash
                addgroup jenkins docker
                docker ps
                '''
            }
        }
        stage('Download ENV') {
            steps {
                withVault([configuration: configuration, vaultSecrets: secrets]) {
                    sh '''
                    docker exec vault sh -c 'export VAULT_ADDR=http://127.0.0.1:8200;rm -rf env.json;vault kv get -format=json kv/dating-app/prod > env.json;exit'
                    rm -rf .env
                    docker cp vault:env.json env.json
                    cat env.json | jq -r '.data.data | to_entries[] | join("=")' > .env
                    '''
                }
            }  
        }
        stage('Build Image') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {
		         sh '''#!/bin/bash
                 docker build -t ubedev/brantas:17 .
                 '''
            }
        }
        stage('Docker Login') {
            steps {
                sh 'echo $DOCKERHUB_CREDS_PSW | docker login -u $DOCKERHUB_CREDS_USR --password-stdin'                
            }
         }
        stage('Docker Push') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            steps {  
                sh 'docker push ubedev/brantas:17'
            }
         }
        stage('Send Discord Notif') {
            when {
                anyOf {
                    expression { return env.GIT_BRANCH == 'origin/master' }
                }
            }
            environment {
                DISCORD_WEBHOOK_URL = credentials('webhook_discord')
            }
            steps {
                discordSend description: "New Dating Service pipeline triggered for $env.GIT_BRANCH", footer: 'Dating Service Pipeline result', link: env.BUILD_URL, result: currentBuild.currentResult, title: JOB_NAME, webhookURL: env.DISCORD_WEBHOOK_URL
            }
        }
   }
    post {
		always {
			sh 'docker logout'
		}
	 }
    }