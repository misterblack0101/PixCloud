package controllers

import (
	"html/template"
	"net/http"
	"pixcloud/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is PixCloud?",
			Answer:   "PixCloud is a cloud storage service for your photos and videos.",
		},
		{
			Question: "How do I sign up?",
			Answer:   "You can sign up by clicking the 'Sign Up' button on the homepage and filling out the registration form.",
		},
		{
			Question: "Is there a free trial available?",
			Answer:   "Yes, PixCloud offers a 30-day free trial for new users.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `You can contact support by emailing <a href="mailto:support@pixcloud.com">support@pixcloud.com</a> or calling our support hotline.`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
