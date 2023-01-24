package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"kubeswitch/internal/cmdutil"
	"kubeswitch/internal/kubeconfig"
	"os"
	"path/filepath"
	"strings"

	"facette.io/natsort"
	"github.com/pkg/errors"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Bye returns a greeting for the given name
func (a *App) Bye(name string) string {
	return fmt.Sprintf("Bye %s, It's nap time!", name)
}

func (a *App) KubeConfig() []string {

	kc := new(kubeconfig.Kubeconfig).WithLoader(kubeconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		if cmdutil.IsNotFoundErr(err) {
			return []string{"kubeconfig file not found"}
		}
		return []string{"kubeconfig error"}
	}

	ctxs := kc.ContextNames()
	natsort.Sort(ctxs)

	// cur := kc.GetCurrentContext()

	return kc.ContextNames()
}

func (a *App) CurrentCtx() string {

	_, err := kubectxPrevCtxFile()
	if err != nil {
		panic("failed to determine state file")
	}

	kc := new(kubeconfig.Kubeconfig).WithLoader(kubeconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		panic("kubeconfig error")
	}

	return kc.GetCurrentContext()
}

func (a *App) KubeSwitch(name string) string {

	prevCtxFile, err := kubectxPrevCtxFile()
	if err != nil {
		panic("failed to determine state file")
	}

	kc := new(kubeconfig.Kubeconfig).WithLoader(kubeconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		panic("kubeconfig error")
	}

	prev := kc.GetCurrentContext()
	if !kc.ContextExists(name) {
		panic("no context exists with the name: " + name)
	}
	if err := kc.ModifyCurrentContext(name); err != nil {
		panic(err)
	}
	if err := kc.Save(); err != nil {
		panic(err)
	}

	if prev != name {
		if err := writeLastContext(prevCtxFile, prev); err != nil {
			panic("failed to save previous context name")
		}
	}

	return strings.Join(listNamespace(), ", ")
}

func writeLastContext(path, value string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrap(err, "failed to create parent directories")
	}
	return ioutil.WriteFile(path, []byte(value), 0644)
}

func kubectxPrevCtxFile() (string, error) {
	home := cmdutil.HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".kube", "kubectx"), nil
}
