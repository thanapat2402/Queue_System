package handler

import "errors"

var (
	ErrService = errors.New("service error")

	MenuFlex = `{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://www.i-pic.info/i/KMdp196143.png",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "Booking",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "How many people are you?",
				  "size": "lg",
				  "color": "#AAAAAA",
				  "flex": 1,
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "Alone",
				"text": "Alone"
			  },
			  "height": "40px",
			  "backgroundColor": "#FBF1C2FF",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "Alone",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "Couple",
				"text": "Couple"
			  },
			  "height": "40px",
			  "backgroundColor": "#DFE9F5",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "Couple",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "Small Group (3-4)",
				"text": "Small Group"
			  },
			  "height": "40px",
			  "backgroundColor": "#F6E6DE",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "Small Group (3-4)",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "The Gang (5-6)",
				"text": "The Gang"
			  },
			  "height": "40px",
			  "backgroundColor": "#D9F0E7",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "The Gang (5-6)",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "spacer"
			}
		  ]
		}
	  }`

)
