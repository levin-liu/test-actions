deploy-dev:
	gcloud functions deploy TestLatency \
      --trigger-event=providers/google.firebase.analytics/eventTypes/event.log \
#      --trigger-resource=projects/mcd-ketchup-dev/events/first_open_2 \
      --runtime=go116 \
      --region=asia-northeast1 \
      --memory=128MB
