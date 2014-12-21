<!-- Title:Vengo 사용하기 -->
<!-- Tags:golang -->

Python을 사용하는 사람이라면, virtualenv을 잘 알것이다. 각 프로그램마다 각자의 파이썬 환경을 구성할수 있는 도구이다. (Py3에 기본으로 추가되어 있다)

go tool은 개인적으로 불만이 많은 도구이다. 특히 go get은 여러가지 기능의 부족으로 당황스러움을 느끼게 한다.

가령 go get은 scm(git,hg...)에서 직접 가져온다. 하지만 언제나 master branch의 최신을 가져온다. 즉 버져닝(versioning)에 대한 부분이 없다.

그리고 go get은 상대경로에 대한 인식을 하지 않는다. 언제나 $GOPATH부터 절대 경로로 패캐지를 임포트해야 한다. (언어의 제안사항이 아님)

go get에 특정 브랜치나 특정 태그(버젼)을 추가하자는 이야기는 go-nuts에 많이 (주기적으로) 나오는 이야기로 알고 있다.

개인적인 생각으로는 특히  오픈소스에서 버젼은 일종의 대화의 좋은 재료라고 생각한다. 구글안에서 처럼 항상 master branch가 stable하다고 하기에는 힘들다.

원래 하려는 이야기는 Versioning이 아니었는데.하튼
[Vengo](github.com/icub3d/goblog)는 Python의 virtualenv처럼 여러 Go환경을 구성할수 있는 프로젝트이다.

![VengoImage1](https://raw.githubusercontent.com/DamnWidget/VenGO/images/vengo.png)

Vengo은 Go 자체도 설치해준다.

``` vengo install -b 1.4 ```

Python virtualenv 에 친숙하다면, 그리 어렵지 않게 사용할수 있을것 같다.

하나 궁금한 점은 이미 만들어진 환경에 go의 버젼을 변경할수 있는지이다.
(다시 만들어도 되겠지만)

Go에서 외부 라이브러리 의존성 관리를 위해 여러 도구들이 있지만. 여러 오픈소스에서 가장 많이 사용하는 도구는 아마 Kz가 만든  [godep](http://github.com/tools/godep) 일것이다.

godep은 현재 패캐지와 의존성이 있는 패캐지들의 정보를 Godeps 디렉토리에  json포맷으로 저장하고 관련 패캐지들을 Godeps/.workspace에 복사한다.

godep을 사용하는 많은 프로젝트가 사용하는 외부 패캐지들을 이렇게 자신의 git 저장소에 저장해서 사용한다.  [GoogleCloudPlatform/kubernetes](https://github.com/GoogleCloudPlatform/kubernetes) , [codeos/etcd](https://github.com/coreos/etcd)

Vengo랑 함께 사용하면 Godep으로 vendoring되어 있는 패캐지들은 현재 환경에 ```godep restore```해서 사용하고 변경된 현재 환경을 ```godep save```해서 사용할수 있다.

요즘 내가 사용하는 방식이다.

godep에서 외부 패캐지를 다 내 저장소에 저장하는 것에 대한 것이 scm을 일종의 디스크로 사용하는 느낌이라서 좋은 방법이 아니지 않나 하는 생각도 든다.

즉 npm이나 rubygem이나 docker registry처럼 scm와 연동된 패캐지 호스팅 서비스등이 있어야 하지 않나 라는 생각인데. Go에서도 비슷한 시도가 있었지만 대중적으로 인기가 있진 않다.
(링크를 찾지 못하겠다.)

오픈소스 커뮤니티에서 여러가지 시도가 있으며, Python virtualenv 처럼 정리되어 go tool에 좋은 것이 포함되었으면 싶다. (*다양성과 효율성*이란 시점에서도 이런 움직임이 있어야 한다고 생각한다.)

- [PackageManagementTools](https://github.com/golang/go/wiki/PackageManagementTools)
