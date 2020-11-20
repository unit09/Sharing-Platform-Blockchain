# Sharing-Platform-Blockchain

## 설명
이 레퍼지토리는 하이퍼레저 패브릭을 이용해 구축한 공유 재산 사용 기록 관리 플랫폼의 블록체인 네트워크와 중간 서버를 구현한 코드를 담고 있습니다.

## 실행 환경
1. Ubuntu 18.04 LTS (가상머신으로 실행)
2. cURL
3. docker 17.06.2-ce 이상
4. docker compose 1.14.0 이상
5. Go 언어 1.11.x
6. Node.js 8.x
7. npm 5.6.0
8. 하이퍼레저 패브릭 1.4
9. mySQL

## 설치 순서
1. 먼저 VMware나 Virtual Box와 같은 가상머신 프로그램을 이용해 Ubuntu 18.04 LTS를 설치합니다.
2. cURL을 설치합니다. 이를 위해 커맨드라인에서 다음과 같이 입력합니다.
~~~
sudo apt install curl
~~~
3. 하나의 가상머신에서 여러 노드들을 실행시키위해 docker와 docker compose를 설치합니다. 설치 및 권한 설정은 아래와 같이 진행합니다.
~~~
curl -fsSL https://get.docker.com/ | sudo sh
sudo usermod -aG docker $USER
sudo reboot

sudo curl -L "https://github.com/docker/compose/release/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
~~~
4. Go 언어를 설치합니다. 설치 및 환경변수 설정은 아래와 같이 진행합니다.
~~~
cd /usr/local
sudo wget https://storage.googleapis.com/golang/go1.11.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.11.1.linux-amd64.tar.gz

echo 'export path=$PATH:/usr/local/go/bin' | sudo tee -a /etc/profile && \
echo 'export GOPATH=$HOME/go' | tee -a $HOME/.bashrc && \
echo 'export PATH=$PATH:$GOPATH/bin:$GOPATH/bin' | tee -a $HOME/.bashrc && \
mkdir -p $HOME/go/{src,pkg,bin}

sudo reboot
~~~
5. Node.js와 npm을 설치합니다.
~~~
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.2/install.sh | bash

sudo reboot

nvm install 8
nvm use 8

npm install npm@5.6.0 -g
~~~
6. mySQL을 설치합니다.
7. mySQL을 이용해 share 유저를 생성하고(패스워드 1124) sharing_platform이라는 DB를 생성합니다. 그리고 해당 DB 안에 auto_increment가 설정된 index(정수형), id(문자열), str(문자열) 3가지 필드를 가지는 identity 테이블을 생성합니다. 
8. 하이퍼레저 패브릭 1.4를 설치합니다. https://github.com/hyperledger/fabric/tree/release-1.4 로 들어가 전체 코드를 다운받은 뒤 $GOPATH/src 에 넣습니다. 
9. 해당 프로젝트를 다운받고 $GOPATH/src에 넣습니다. 이것으로 실행을 위한 사전준비가 끝납니다. 각 단계별로 설치가 끝날때마다 프로그램 별 버전확인 명령어를 입력해 설치가 제대로 되었는지 확인하는 것이 좋습니다.

## 실행
1. basic-network 안에 있는 start_channel.sh를 실행합니다. 이 쉘코드는 각 노드를 도커에 올리고 채널을 만든 후 각 노드를 채널에 가입시키는 과정을 포함합니다.
2. basic-network 안에 있는 chaincode.sh를 실행합니다. 이 쉘코드는 체인코드를 채널에 설치하고 인스턴스화하는 과정을 포함합니다.
3. basic-network 안에 있는 start_ca.sh를 실행합니다. 이 쉘코드는 외부 접속을 위한 인증서를 관리하는 CA를 실행하는 과정을 포함합니다.
4. application/rest 안에 있는 server.js를 실행합니다. Node 서버를 실행하는 것이기 때문에 사전에 application 안의 package.json에 포함된 패키지들을 설치하여야합니다. 실행 명령어는 node server.js입니다. 또한 server.js를 열고 HOST를 서버가 실행되는 ip 주소 혹은 localhost로 변경해줍니다. 그리고 connection 객체의 설정을 적절하게 바꿉니다. 만약 위의 설치순서 7번에 맞춰 유저와 데이터베이스를 생성하였을 경우 이 부분을 건드리지 않아도 됩니다.
5. 이제 REST API의 GET 방식을 통해 서버가 잘 작동하는지 테스트할 수 있습니다. 자세한 테스트 내용은 하단의 테스트 예시를 참고하십시오.
6. 모든 테스트가 끝난 후 서버를 종료한 후 basic-network 안에 있는 kill.sh를 실행하여 모든 네트워크를 종료합니다.

## 테스트 예시

1. 사용 시작 기록

이 기능을 사용하면 블록체인에 대여의 시작에 대한 기록을 저장합니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직), id(유저의 아이디), target(빌리고자 하는 대상 아이디), location(대여 시작 위치)가 있습니다.

![Screenshot from 2020-11-20 23-54-05](https://user-images.githubusercontent.com/45275382/99815214-5ddbc880-2b8d-11eb-8b18-df0e4a8aa5c0.png)

이 프로젝트에서는 체인코드에서 각 조직별로 할 수 있는 기능을 제한함으로써 접근 제어를 구현합니다. 예를 들어 이 기능은 서비스 제공 기업(share1)은 사용가능하지만 지방자치단체(agency)의 경우에는 이 기능을 사용할 수 없기 때문에 해당 기능을 실행하는 조직을 agency로 하여 요청을 하면 아래와 같은 결과가 나타납니다.

![Screenshot from 2020-11-20 23-57-09](https://user-images.githubusercontent.com/45275382/99815227-603e2280-2b8d-11eb-96f5-4f45852b84a6.png)

2. 사용 종료 기록

이 기능을 사용하면 블록체인에 대여의 종료에 대한 기록을 저장합니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직), id(유저의 아이디), target(빌리고자 하는 대상 아이디), location(대여 종료 위치)가 있습니다.

![Screenshot from 2020-11-20 23-54-20](https://user-images.githubusercontent.com/45275382/99815218-5f0cf580-2b8d-11eb-8c26-b9687a8ea134.png)

3. 전체 사용 기록 조회

이 기능을 사용하면 블록체인에 기록된 모든 사용 기록을 가져옵니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직)가 있습니다.

![Screenshot from 2020-11-20 23-55-07](https://user-images.githubusercontent.com/45275382/99815219-5fa58c00-2b8d-11eb-9f49-7d01856863f4.png)

위의 결과처럼 id에는 랜덤한 식별코드가 들어가기 때문에 블록체인의 기록만 가지고는 누구의 기록인지 알 수 없습니다.

4. 위치로 사용 기록 조회

이 기능을 사용하면 블록체인에 기록된 특정 장소에 대한 모든 사용 기록을 가져옵니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직), location(조회할 위치)가 있습니다.

![Screenshot from 2020-11-20 23-57-43](https://user-images.githubusercontent.com/45275382/99815229-60d6b900-2b8d-11eb-9f70-b3c250dd8102.png)

5. 장소별 사용 시작/종료 횟수 조회

이 기능을 사용하면 블록체인에 기록된 모든 사용 기록에 대하여 각 장소별 횟수를 계산하여 보여줍니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직)가 있습니다.

![Screenshot from 2020-11-20 23-58-16](https://user-images.githubusercontent.com/45275382/99815230-60d6b900-2b8d-11eb-9785-053f769df44d.png)

6. 특정 유저의 사용 기록 조회

이 기능을 사용하면 특정 유저에 대하여 블록체인에 기록된 모든 사용 기록을 가져옵니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직), id(조회하고자 하는 유저의 아이디)가 있습니다.
이때 id를 이용하여 별도 데이터베이스를 조회하여 식별코드를 받아온 후 이 식별코드들을 이용해서 조회를 합니다.

![Screenshot from 2020-11-21 00-04-07](https://user-images.githubusercontent.com/45275382/99815234-616f4f80-2b8d-11eb-9b12-17013e3b8b62.png)

![Screenshot from 2020-11-20 23-59-27](https://user-images.githubusercontent.com/45275382/99815232-616f4f80-2b8d-11eb-9438-0f7f6315ea64.png)

7. 거치대 위치 등록

이 기능을 사용하면 블록체인에 거치대의 위치를 기록합니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직), id(거치대 아이디), location(거치대 위치)가 있습니다.

![Screenshot from 2020-11-20 23-55-42](https://user-images.githubusercontent.com/45275382/99815221-5fa58c00-2b8d-11eb-9963-423cc186cdf3.png)

8. 거치대 위치 조회

이 기능을 사용하면 블록체인에 기록된 모든 거치대의 정보를 가져옵니다.
이 기능을 사용하기 위한 인자로는 ccp(해당 기능을 실행하는 조직)가 있습니다.

![Screenshot from 2020-11-20 23-56-31](https://user-images.githubusercontent.com/45275382/99815224-603e2280-2b8d-11eb-92d4-2cb55f29e65f.png)
