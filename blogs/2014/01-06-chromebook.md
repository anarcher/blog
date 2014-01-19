<!-- Title: Chromebook for coding -->
<!-- Tags: chromeos -->
<!-- Created: 2014-01-06 -->
<!-- Updated: 2014-01-06 -->


![enter image description here][1]

[크롬북](http://www.theverge.com/2013/10/23/4948120/acer-c720-chromebook-review)을 샀다. 

크롬북의 가장 큰 장점은 **가격**이다. $199에 이정도 노트북을 살수 있다는 것은 분명 매력적.
정말로 크롬 웹 브라우저만 있다. 흠... 

- 좋지 않은 디스플레이 (흐릿한 디스플레이는 C720의 가장 큰 하드웨어 단점이다)  
- 8시간 사용 가능한 배터리 (하스웰 기반 CPU이 주는 혜택. 8시간이상 되니까 핸드폰처럼 충전기를 가지고 다니지 않게 되었다)
- 내 아이폰과 같은 용량인 16기가 SSD.

나의 개발환경을 기본적으로 한쪽 화면에 터미널(들),또 다른 화면에는 웹 브라우저(들)뿐이다. 크롬OS는 개발자 모드로 shell을 사용할수 있으므로, 코딩하는게 크게 문제는 없다. 

![enter image description here][2]

Acer C720의 개발자 모드 사용법은 [http://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices/acer-c720-chromebook](http://www.chromium.org/chromium-os/developer-information-for-chrome-os-devices/acer-c720-chromebook) 참고.

Ctrl + alt + t 을 누르면 crosh가 뜬다. 하지만 크롬탭으로 실행되기때문에 단축키 충돌등이 발생한다. 

[crosh window](https://chrome.google.com/webstore/detail/crosh-window/nhbmpbdladcchdhkemlojfjdknjadhmh) 을 설치하면 window으로 따로 띠우기 때문에 사용하기 좀 더 편하다. (문제는 터미널을 탭으로 더 띠울수 없다)

그리고 Ubuntu/Debian을 사용하기 위해서 [crouton](https://github.com/dnschneid/crouton)을 설치했다. 
(cli-extra으로 CLI만 사용하도록 설치)

거의 회사에서 쓰는 ubuntu와 비슷하다. 단 다음과 같은 사용시에 단점들이 있긴 하지만.

- crosh window에서 한글이 되지 않는다. 
- docker가 설치되지 않는다. chromeos의 linux kernel이 docker가 필요하게 구성되어 있지 않다. 

싸고 터미널과 웹만 있으면 코딩이 가능한 사람들에게는 추천할 만하다고 생각한다. 


  [1]: https://fbcdn-sphotos-b-a.akamaihd.net/hphotos-ak-ash4/1509213_10152197112323647_1918159647_n.jpg
  [2]: https://fbcdn-sphotos-h-a.akamaihd.net/hphotos-ak-prn2/t1/1517671_10152197112378647_1643871268_n.jpg
