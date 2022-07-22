package service

var (
	QueueFlex = `{
		"type": "bubble",
		"size": "kilo",
		"direction": "ltr",
		"hero": {
		  "type": "image",
		  "url": "https://www.i-pic.info/i/KMdp196143.png",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "position": "relative"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "md",
		  "contents": [
			{
			  "type": "text",
			  "text": "%v",
			  "weight": "bold",
			  "size": "xl",
			  "gravity": "center",
			  "margin": "lg",
			  "wrap": true,
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "Name",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "margin": "xs",
				  "contents": [
					{
					  "type": "text",
					  "text": "Date",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
					  "contents": []
					}
				  ]
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "margin": "xs",
			  "contents": [
				{
				  "type": "text",
				  "text": "Queue",
				  "size": "sm",
				  "color": "#AAAAAA",
				  "flex": 2,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "%v",
				  "size": "sm",
				  "color": "#666666",
				  "flex": 4,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "spacer",
				  "size": "xs"
				},
				{
				  "type": "image",
				  "url": "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=%s",
				  "size": "md",
				  "aspectMode": "cover"
				},
				{
				  "type": "text",
				  "text": "You can enter the restaurant by using this code instead of a ticket",
				  "size": "xxs",
				  "color": "#AAAAAA",
				  "margin": "xxl",
				  "wrap": true,
				  "contents": []
				}
			  ]
			}
		  ]
		}
	  }`
)
