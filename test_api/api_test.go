package test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}
	dir, _ := ioutil.TempDir("", "dockertest")
	defer os.RemoveAll(dir)
	dockerfilePath := dir + "/Dockerfile"
	dockerfile := `FROM golang:1.19-alpine
	WORKDIR /src
	COPY $HOME/go-api-http-test
	RUN go build $HOME/go-api-http-test/cmd/api	
	CMD ["./api"]`
	ioutil.WriteFile(dockerfilePath,
		[]byte((dockerfile)),
		0o644,
	)

	appDocker, err := pool.BuildAndRunWithBuildOptions(
		&dockertest.BuildOptions{
			ContextDir: dir,
			Dockerfile: "Dockerfile",
		},
		&dockertest.RunOptions{
			Name: "buildarg-test",
		}, func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	appDocker.Expire(60)
	appHostAndPort := appDocker.GetHostPort("8888/tcp")
	log.Println("start app on host: ", appHostAndPort)
	code := m.Run()
	if err := pool.Purge(appDocker); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

/*
func TestApi_EntityRead(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Тест API. Метод Read",
			fields: fields{
				ctx: ctxTest,
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := newInfoblocksDAO(tt.fields.ctx)
			_, err := d.read(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/
