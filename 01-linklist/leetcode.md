@[TOC]

## 链表

### leetcode 237:删除链表中的节点

1. 题目

    请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
    
    现有一个链表 -- head = [4,5,1,9]，它可以表示为:
    
    4->5->1->9

    示例 1:
    
    输入: head = [4,5,1,9], node = 5
    输出: [4,1,9]
    解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
    示例 2:
    
    输入: head = [4,5,1,9], node = 1
    输出: [4,5,9]
    解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
    
    说明:
    
    链表至少包含两个节点。
    链表中所有节点的值都是唯一的。
    给定的节点为非末尾节点并且一定是链表中的一个有效节点。
    不要从你的函数中返回任何结果。

2. 解题思路
    
  一般来说，我们要删除单链表的一个节点，我们需要找到需要删除节点的前一个节点，然后用这前一个节点直接指向需要删除节点的后一个节点，直接跳过要删除的节点，达到删除的效果。

 


![avatar](data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAA5sAAAA7CAYAAAD1n55jAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAABGgSURBVHhe7Z0JjF3TH8er1tiqRKmq2GprWkJqKFK0RG1VMukgRFBRrUhIyCgVgkpDY6mK0ib2tukWe621tQRRQmqPJTGKIrUbnH++xzvP6zvnf3s6nc78bvv5JL905s6beb93+znnnt89597bxQEAAAAAAAC0MxSbAAAAAAAA0O5QbAIAAAAAAEC7Q7EJAAAAAAAA7Q7FJgAAAAAAALQ7FJsAAOBpaWlxTU1NrrGxkSDaFM3NzRWbAAAAKDYBAKDCAw884Lp06ZIsIoj/QvuooaEh+bN1ObRfFAAAAAGOCgAA4AnFJhSjfaR9BSuCPwAAUA9HBQAA8FAs5EGxmQZ/AACgHo4KAADgoVjIg2IzDf4AAEA9HBUAAMBDsZAHxWYa/AEAgHo4KgAAgIdiIQ+KzTT4AwAA9XBUAAAAD8VCHhSbafAHAADq4agAAAAeioU8KDbT4A8AANTDUQEAADxWioWlS5e6K664wueieOihhyo/sQHFZhqKTQAAqIejAgAAeCwUCyo0Fy1aVPnO+UJTOd14442VLZ0PxWYaik0AAKiHowIAAHgsFAu1hWZAOVkqYig201BsAgBAPRwVAADAY7FY+PHHH31OWlZrBYrNNBSbAABQD0cFAADwWCsWPv/88+q1mx988EFla+dDsZmGYhMAAOrhqAAAAB5LxYIKTeUSgms27UOxCQAA9XBUAAAAj8ViYfHixdXZzSlTplS2di4Um2koNgEAoJ7oqNDS0uKamppcY2MjQbQpmpubKzaVD/wnVjfK7L/VYkFLaJWXldwoNtOUvdik/ydWNxj/EOty/D//o6NCOFik/gjxX2gfNTQ0JH+2Lof2S5kHG/ifF/ifjrXFf4tY2rfKg2IzpuzFJv1/Xmgf0f/Hof2C/2t/4H86ivyPtpb9YNFRaB8x2IhZWwYbUAz+p7Hiz/Lly912223nTjzxRDd16lT36aefVn5SjFX/wx1p9cxNC+B/mrL3n2XPv6PA/zT4v26A/2mK/Im2IlseyJam7P7gfx74n8aKP1oOpTwU66+/vv+3V69e7uyzz3b333+/++qrryqvXBEL+Z9wwgn+ZkC6QZBQoalrNnn0iX3K3n+WPf+OAv/T4P+6Af6nKfIn2opseSBbmrL7g/954H8aK/60trb6POpjww03dF27dvVf77777m706NFu9uzZ7rvvvvO/ZyH/hx9+eIWcVXguWrSo8lMbKC/8j7HUf86aNcvnohMs77zzTmVrMZbytwz+pym7P/ifB/6nKfIn2opseSBbmrL7g/954H8aS/5069bN51IUG220kVtvvfV89O3b1x100EF+OxSjfYT/McH/p59+2hd4mkHXiY/OYOLEiT4XnWCR34MHD3bz5893//zzT+UVMfT/eeB/mrL7g/954H+aIn+irciWB7KlKbs/+J8H/qex4M/vv//uvvzyS9e7d2+fS1sCitE+wv+Y4H99bLPNNm7vvfd2hx12mDv55JPdBRdc4MaNG+duu+02N2PGDLdgwQL37rvvuqVLlxYWg6vCZZdd5jbeeONqDhtssIH/d6+99nLTpk3z7aQe+v88tI/wP6bs/uB/HvifpsifaCuy5YFsacruD/7ngf9p1oQ/mhnSDNHbb7/tZ4z0HjfffLMbO3asO/fcc92wYcPcwIED3R577OG22GIL//6KHXbYofp1UYRBeL9+/dyoUaP811CM9hH+xwT/tSx7yZIl7sUXX/TLtG+//XZ39dVXuzFjxrgRI0a4QYMG+Zn0Hj16VJd1h9D3urmVfn744Yf711944YX+9ydPnuzmzJnjXnrpJf/3ly1bVnnnGC2fDW7X/32FCuDrrrvOffvtt5XfoP/PRfsI/2PK7g/+54H/aYr8ibYiWx7Ilqbs/uB/HvifJsefv//+28/gaCZHMzrTp0/3Mzya6dGMj2Z+NAOkGRgNiPX3akPLAnv27On69+/vhgwZ4k477TR30UUXuWuvvdbdddddbt68ee6VV15xxx13nF8+WP/7IcLywqFDh7rnnnvO54b/eWgf4X9MW/xRe/j66699e3j++ed9e7j11lt9e9DJj+HDh1fbw9Zbb131N4Q81okVtYejjjrKnX766b49DBgwIHptfagY1ezn+eef75/liv95aB/hf0zZ/cH/PPA/TZE/0VYrsukOhLopxJQpU/zdCa2BbGnK3llZyV934tQASLno32effbbyExvgf5rgj2ZgNBOjGRnNzGiGRjM1mrFJzeRohkc/14yPXq8ZIP2+ZoQ0k6MZIs3khBv55KCBugbite+l0ABb12qOHDnSvf/++5VX/0vZ229HoX2E/zEd4c+ff/5Znel/6qmn/N2Vw0z/Oeec4x/3o5n+3Jl9RTjxonao762hcZClvJQL/seUvf+0lP/ixYt9LiE0DrKC8sH/mCJ/oq1WZAu3ug+iWQPZ0ljqrNqChfx1okV35Axf69mCyilsswD+pwn+aEZSMzGakdHMjAq/cI2aZm40gxOuUdPMzprgqquuql6zpoG0itru3bv7IrZ26WAtZW+/HQX+p7Hkjx71o1xyon4FgCVqB91WUC74H1P2/tNS/uEESwjGP/Yp8ifaaq2xBNGsgWxpLPmjJVV6xuCECROqz+xbGRbyT3Wq1toB/qex5P+kSZOq3vTp08fdfffdyZui1GIpf8vgfxpL/tTeHCgVmt3Xv3rdEUcc4caPH+9PxGibFcIzZkPOVlAu+B9jyf+yjn8ClorLevA/TZE/0VZLsgnlYimfALKlseRPuFmKlg3qzLUe66Blid98803lFTHW/A8oJ5aR2MeSP1oie/TRR7snnngi+w6fVv23Bv6nseLPr7/+6vOojVBc6njQ0NDgZ/5feOEF98cff1R+y57/esasVj+Ez2AF5YL/MZb8KfP4R8Wx8tCJFmvPWBb4n6bIn2irpcYilIulfALIlsaSP1rGGPxRaBmhOl6d7dMg/N5773XLly+vvPpfrPkvdHZbObGMxD4W/VkVyp5/R4H/aaz488UXX/g8FOr399tvP/8olCeffNL9/PPPlVfFWPJf1+mHgXb4LFZQLvgfY8mfMo9/NNapzV33bdFJFysoJ/yPKfIn2mqpsYggmzWQLY0lf+o729pQh6vOV2e7TznlFDd37ly/xNCa/0KDDnW2KjqtoH2E/zEW/VkVyp5/R4H/aaz489dff7mdd97ZD6ofeeQR/30OVvLXwFrXrAWUk6V2if9pLPWfZR//aLyj65XDMvLa9tDZKB/8jynyJ9pqSTahXCzlE0C2NMGfmTNndnrssssuVX+KIiwz2Xzzzf2Dx7XNEio0rS0l0T7C/5ha/x977DH/XMwQb7zxhg8tb9XdNH/55ZfKb9lBg47QLgiirdHZaNl4bT7dunVzZ511lp8x+e233yqvirEy/qkfWIfPYQXlQv8fw/hnzaD2YOmpFNpH+B9T1H9GW610tgHlYlF+ZEsT/LEQqWcU5oYVdCdaS2f0AtpH+B+zqv7rkQvydLfddnMHH3ywO+mkk9zo0aPdNddc42/oo4L1vffe89egdQStra3+IfypwQtB5MTChQsrNnUuqbamQfUmm2zin2WrtvrDDz9UXv0vof12JiqI62/oEj6DFZQL/X/Mqvb/azLWhvFPIFxKZAXlgv8xRf1ntNVCZ1uLcrGUTwDZ0ljyp2gZSYhwVm/TTTd1Z5xxhrv00kv9dguEJSQW0T7C/5ha/3WA/P77733oUSOffPKJD/2/6sYkGlTqGYG6acP111/vH0Tf1NTkH5ey5557us0226zqqULP6NRNHk499VTX3Nzs7rzzTjd//nz/MPqV3WUWYF2jvv3Uhvr9cA3bkUce6e644w6/2sDC8as+1/qwgPKg/49h/LPm4AaJ9inyP9pqqbEI5WJRfmRLU4bONlyvoDPdegD4rFmzqjNHVvLXNTu6E2EtKlKsdLjaR/gf097+6AYOenj9vHnz/IPrVZAOGzbM7bvvvn5pYHBaA4Ydd9zRHXrooX7QcOWVV7qpU6e6xx9/3P++BtK5160BlBHNUurEzrJly/xJnaJiszZ0LNAxoXabNazlpVzo/2MY/6wZdOJW966wgvYR/scU+RNttSRbmDpX6GtLIFsaS/7UdrYajIez2Xqm2rRp05JOWchfhaauTwi514aVO9IqF/yP6Wh/NLjWdaAaMOjkxJgxY9zxxx/v+vbt689W17oj97fffnvXr18/N3jwYDd8+HB/HZsK2HHjxrmbbrrJL93VUshwvelrr73m3nzzTT9419I+vZ/Fa02hfdF1jWFWXo9KCLPyCvkQrj9esGBB9ZrkRx991Lszffp0P2OvmXeFTpLccMMNPrRSQ3eFVYwaNcqdd955PhobG31oeeuQIUN8DBo0yB1wwAE+5POuu+7qQw53797dh/r0WsdrY2XP2awNHR9qv7eGtbyUC/1/DOOf1UeXDtUWljruWBn3BPA/TZE/0VYrjUU5pMIKygXZYqz4I4Iz6mg1YLnllltcS0tL5adpLOSv2cuQe31oyaQFlAv+x1jyX2hAsWTJEr9sV0WA2sDYsWP9AH/EiBHumGOO8deK7rPPPq5Xr17VZ7PlhIpZDfh32mknXwTsv//+bsCAAb5QGDp0qC8ezjzzTP9eF198sS8w9GzDUHhMnjzZFyOhwFVoUBGKFxW6KmjeeuutaqGjGdpQBCmsnITUMubavPTojZDzxx9/XC3OFLqmMXxGhWatw+dXhCJNMWnSpOr+UmjGOhRr2qehWBs5cmS1YFPo/zUUbYcccki1aOvfv3+1aFOEom1lhduqhv5W+Lta/h3eT0VjyEXFZMhRRWbIPXwmFaPhs6pIDftAxWvYPypqw37T82TDPn399dddjx49krnVhmZ39K8G5hMnTvTLafW9NUK+VlAu9P8xjH9WHx0DQu5q91rRZQ3lhv8xRf5EWy01FssgWxpL/mgwooGKBny54H8e+J9mbfFHy3dVNKntfPjhhysUSRoMaHCvs+NqYxMmTHDjx4/3be2SSy7xhYKKTBUOxx57rC8mDjzwQD/g0aC+vtDRXRC1z9o7tMw4vEdt9O7du5qDQjfSSL2uKLbaaqvke7ZHbLnlltX3UW61ueqkQCjWtE9DsabQIwxCwaYZ61C0abY7FG261re2eA0Fv+K+++6rFm4qgEPh9swzz6xQLMuHUEjXFv8//fRTxR4b9OnTJ7l/u3bt6kPLbFWkv/rqq5XfoP/PRfuI/j/Gkj+Mf9Yc+J+myJ9oK7LlgWxpyu4P/ueB/2nwZ/XQA/dD8RIKGs3MhkJHM7ShCFKEpZv18eCDD1aLqPqoLbZCpF63spgxY0b1/ebMmbNCXi+//PIKBdpHH31U/TyfffZZ9TMqih7FAW1Hs6hqiyHC3WgHDhzo7rnnnuQdnmm/eWgf0f/HlN0f/M8D/9MU+RNtRbY8kC1N2f3B/zzwPw3+ANigZ8+evi0qtt12W3f55Zf7Jc1F0H7zoP9PU3Z/8D8P/E9T5E+0FdnyQLY0ZfcH//PA/zT4A2AD3fRKy451s6vcOzHTfvOg/09Tdn/wPw/8T1PkT7QV2fJAtjRl9wf/88D/NPgDUF5ov3nQ/6cpuz/4nwf+pynyJ9qKbHkgW5qy+4P/eeB/GvwBKC+03zzo/9OU3R/8zwP/0xT5E21FtjyQLU3Z/cH/PPA/Df4AlBfabx70/2nK7g/+54H/aYr8ibYiWx7Ilqbs/uB/HvifBn8AygvtNw/6/zRl9wf/88D/NEX+RFuRLQ9kS1N2f/A/D/xPgz8A5YX2mwf9f5qy+4P/eeB/miJ/oq3IlgeypSm7P/ifB/6nwR+A8kL7zYP+P03Z/cH/PPA/TZE/0da5c+f6FxPE6kRZwX+iPQIAygf9P9EeUVbwn2iPSBFtbW1tdbNnz3YzZ84kiDbFwoULKzaVD/wnVjfK7D/Augz9P7G6wfiHWJfj//nPKXgAAAAAAABodyg2AQAAAAAAoN2h2AQAAAAAAIB2h2ITAAAAAAAA2h2KTQAAAAAAAGhnnPsfq2/jIlEMQu0AAAAASUVORK5CYII=)


  
  但是这一题并没有给我们头结点，我们无法通过遍历的方式来找到这前一个节点。这题只给了我们

3. 代码实现
 
       /**
        * Definition for singly-linked list.
        * type ListNode struct {
        *     Val int
        *     Next *ListNode
        * }
        */
         
       func deleteNode(node *ListNode) {
           node.Val=node.Next.Val
           node.Next=node.Next.Next
       }