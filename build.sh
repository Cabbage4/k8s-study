rm main
export GOOS=linux && export GOARCH=amd64 && go build -o main .
docker build -t cabbage4/go-demo:v0.0.1 .
docker push cabbage4/go-demo:v0.0.1
rm main
