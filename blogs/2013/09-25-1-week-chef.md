<!-- Title:Chef 사용기 -->
<!-- Tags: Chef,Devops,ServerConfiguration -->

> There are only two hard things in Computer Science: cache invalidation and naming things.
> -- Phil Karlton 

Chef을 써보면서, 요리에 관련된 메타포를 사용하여 Chef 구성요소들의 이름들을 잘 지었다는 생각이 든다. 
Chef을 구성하는 요소에 대한 설명을 간단히 하자면:

- Chef    : 요리사,Server Configuration Project 
- Recipe  : 요리법,어떤 작업을 할지 서술하는 목록 
- Cookbook: 요리책,Recipe의 모음집.
- Knife   : 칼 요리도구,Chef의 CLI 도구
- Node    : 구성 작업할 서버
- Workstation : 작업장, Knife등의 도구를 통해 Chef에 작업을 하는 머신. 
- Attribute : 
- Data bag : 
- Envoiroment: 
- Role :

물론 모든 단어가 요리에 관련된 단어는 아니다. 좀더 명확하게 설명할수 있는 요소는 굳이 사용하지 않은 듯 하다. 

다른 좋았던 점은 Data bag과 Search 부분이다. 구성에 대한 정보를 Data bag의 Attribute으로 등록해두고 Recipe에서 사용하는 방식은 기능과 정보를 나누어서 관리 할수 있어서 데이터에 따른 recipe 사용이 좀더 편해진다. 

특히 Chef REST API을 가지고 서버 구성하는 애플리케이션을 만들때, 따로 DB을 구성하지 않고 Chef의 Data bag을 사용할 수 있다. 

구성이 끝나면,Chef의 Node에 추가되는데. Node의 정보(cpus,mem...)등도 함께 등록되기 때문에 현재 가용중인 서버들의 정보를 얻기에 쉽다. 

이 부분에서 좀 아쉬운 부분은 구성 작업이 성공했는지 실패했는지에 대한 정보가 Node에 없어서 관련 Recipe을 만들어 줄 필요가 있다. 

그리고 Chef 문서에도 있는 데, Databag이나 Cookbook등을 SCM(버젼관리) 저장소를 쓰는 것을 추천하는데. Chef server에 git이나 hg등의 scm 기능이 추가되어도 괜찮지 않을까 생각이 든다. 

또 다른 아쉬운 점은 Chef은 기본적으로 하나의 노드에 대한 작업에 대한 부분만 고려되어 있는 듯 한 느낌이다. 즉 여러 노드를 병렬/순차적으로 작업해야 할때에 대한 기능이 직접적인 지원이 없다. 
(예를 들어 AWS의 Cloudformation 같은 기능을 말한다)

자체적으로 데이터 스토어가 있다는 점, 중앙 관리가 가능하다는 점에서는 Chef 사용에 있어 좋은 점이라고 생각된다. 

Chef을 사용하지 않고 하려면 어떤 도구들을 사용하면 될까? 
개인적으로는 Ansible와 Zookeeper을 가지고 비슷하게 구성해볼수 있지 않을까 생각된다. 

