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

	ReportFlex = `{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://www.i-pic.info/i/KMdp196143.png",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"label": "Action",
			"uri": "https://linecorp.com"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "md",
		  "action": {
			"type": "uri",
			"label": "Action",
			"uri": "https://linecorp.com"
		  },
		  "contents": [
			{
			  "type": "text",
			  "text": "Queue Status",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "contents": [
					{
					  "type": "text",
					  "text": "Type",
					  "weight": "bold",
					  "margin": "sm",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "Current",
					  "weight": "bold",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "Waiting",
					  "weight": "bold",
					  "color": "#AAAAAA",
					  "align": "end",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "contents": [
					{
					  "type": "icon",
					  "url": "https://cdn-icons-png.flaticon.com/512/32/32384.png"
					},
					{
					  "type": "text",
					  "text": "A",
					  "margin": "lg",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "align": "end",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "contents": [
					{
					  "type": "icon",
					  "url": "https://cdn-icons-png.flaticon.com/512/32/32384.png"
					},
					{
					  "type": "text",
					  "text": "B",
					  "margin": "lg",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "align": "end",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "contents": [
					{
					  "type": "icon",
					  "url": "https://cdn-icons-png.flaticon.com/512/33/33308.png"
					},
					{
					  "type": "text",
					  "text": "C",
					  "weight": "regular",
					  "margin": "lg",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "align": "end",
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "contents": [
					{
					  "type": "icon",
					  "url": "https://cdn-icons-png.flaticon.com/512/32/32441.png"
					},
					{
					  "type": "text",
					  "text": "D",
					  "weight": "regular",
					  "margin": "lg",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%v",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "align": "end",
					  "contents": []
					}
				  ]
				}
			  ]
			},
			{
			  "type": "text",
			  "text": "Wear masks and maintain safe social distancing limits during your visit",
			  "size": "xxs",
			  "color": "#AAAAAA",
			  "wrap": true,
			  "contents": []
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "Check my queue",
				"text": "ตรวจสอบคิวของฉัน"
			  },
			  "color": "#49C09AFF",
			  "style": "primary"
			}
		  ]
		}
	  }`
)
