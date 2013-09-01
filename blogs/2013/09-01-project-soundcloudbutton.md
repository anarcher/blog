<!-- Title:기분전환용 프로젝트 Soundcloud control button -->
<!-- Tags: myproject,chrome -->

[@Outsider](https://twitter.com/outsideris)님의 [기분전환용 프로젝트 Gittip-links](http://blog.outsider.ne.kr/973)에 영감을 받아서 나도 한번 만들어 봤다. 

# 동기

[Soundcloud](https://soundcloud.com/)에서 이런 저런 Podcast을 자주 듣는 편이다. 주로 크롬 웹 브라우저에서 듣는데. 보통 탭이 한두개가 아닌지라 Soundcloud에서 듣다가 음악이나 방송을 끄거나 하려면 종종 열려 있는 탭에서 찾아야 하는 점이 불편했다. 그래서 크롬 익스텐션으로 만들어서 쉽게 (혹은 빠르게) 재생/중지를 할수 있으면 좋지 않을까 싶었다. 

만들고 나서, [soundcloud button extension](https://www.google.co.kr/search?q=sound+cloud+button+&oq=sound+cloud+button+&aqs=chrome..69i57j69i60j0l2j69i60l2.5935j0&sourceid=chrome&ie=UTF-8#newwindow=1&psj=1&q=soundcloud+button+extension) 으로 구글링해보면, opera용 add-on 이 있기는 하다. (사실 설명 부분은 살짝 베끼긴 했다. 영어 실력으로 부족으로.;ㅁ;)


# 개발

사실 크롬 익스텐션을 만들어 본적이 없었기 때문에, 2시간-프로젝트(2HourProject)가 될거라고는 생각하지 않았다. 그래도 https://github.com/outsideris/gittip-links 코드도 좀 보고,구글 튜토리얼도 보면서 이것저것 HelloWorld 수준의 코드를 만들어 보면서 만들긴 했다. 대충 업로드와 이미지 작업 빼고 5시간 정도 걸린듯 하다.(카페에서 친구와 잡담도 하고 멍하니 있기도 해서 정확한 시간은 모르겠다.;) 

하면서 아이콘 이미지 만드는 것 하나 잘 모른다는 걸 알았다. Soundcloud의 버튼을 쓰고 싶었지만, 
배경을 투명하게 만드는 것도 하기 힘들어서 포기하고 http://findicons.com/icon/248309/play 을 가져다가 썼다. (라이센스는 프리웨어,비상업)

이미지 툴도 없어서 GIMP을 맥에 설치할까 했지만, 웹에 있는 http://pixlr.com/editor/ 을 알게 되어서 조금 써볼려고 해봤다. (결론은 실패.:-( )

크롬 익스텐션의 구조는 매우 간단해서 background page에 soundcloud page에 반응하는 content script가 dom을 읽어서 Message passing 하는 구조이다. 

크롬 익스텐션이 의외로 어렵진 않았는데, 간단한 것을 만들어서 크게 내부 구조를 알 필요가 없긴 했다. 


# 에필로그 

나 혼자 쓰게 될것 같은데. ㅎㅎ 우선 간단한 아이디어라도 구현해보니 나쁘지 않았다. 
만들자 마자 올려서 어떤 버그가 있는 지도 안알라줌. (나도 모르지만 존재 여부는 당연히.. ;ㅁ;)

소스는 [여기](https://github.com/anarcher/SoundCloudControlButton) 에 있고,익스텐션은 [여기](https://chrome.google.com/webstore/detail/soundcloud-control-button/ncmhcpmbfcnmnbhfpndnfmjgifcifamn)에 있다. 

ps)재미삼아 오마주한 프로젝트의 블로그 글을 따라했다. ㅎㅎㅎ


