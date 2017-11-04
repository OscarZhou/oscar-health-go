# oscar-health-go
An E-commercial website developed by Golang

# OAuth2 

I used the OAuth2 to implement the login part of this website. Honestly, I was very happy to learn the related knowledge. If my boss did not tell me about OAuth2 standard. I probably won't know it. There are four methods to finish the authorization:  

**Authorization Code**  

Authorization Code is the most complex method among these four methods. There are several steps to be finsihed.  

1. The client requests the grant code from the resource owner. At the stage, the content we need to send is:  

		{
			"response_type": 	"code"
			"client_id":		""
			"redirect_uri":		""
			"scope":			""	//Optional
			"state":			""	//Optional
		}	
		
If you want to make this process safer, you can assign `state` a random value and then send it. You can validate the `state` returned from the resource owner. We need to use `GET` to request it. And then the resource owner accept your request, it will response you two parameters by url, which are  

		{
			"state":			""
			"code":				""
		}
		
2. When we get the returned values, we need to send the grant code to the authorization server, the content of the stage we need to send is:  

		{
			"grant_type":	"authorization_code"
			"code":			""
			"client_id":	""
			"redirect_uri":	""
		}

As the current authorization method is authorization code mode, so the field of `grant_type` is fixed. The information here needs to be sent by `POST` method. At the end of this stage, the authorization server returned some values, like  

		{
			"access_token":		""
			"token_type":		""
			"expires_in":		""
			"refresh_token":	""
			"scope":			""
		}

The token will be used to send the resource server to request the protected resources.  

3.		

