package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// only works if no other image has been created in the meantime
func tagDockerImage() error {
	cmd := exec.Command("docker", "load", "-i", "dagger-hello-deno.tar")
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("sh", "-c", "docker images -q | head -n 1")
	imageID, err := cmd.Output()
	if err != nil {
		return err
	}

	cmd = exec.Command("docker", "tag", strings.TrimSpace(string(imageID)), "dagger-hello-deno:latest")
	return cmd.Run()
}

func build(ctx context.Context) error {
	// Initialize the Dagger client
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Use Ubuntu as the base for compilation
	ubuntu := client.Container().From("ubuntu:22.04")

	// Install Deno
	deno, err := ubuntu.WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "curl", "unzip"}).
		WithExec([]string{"sh", "-c", "curl -fsSL https://deno.land/x/install/install.sh | sh"}).
		WithExec([]string{"mv", "/root/.deno/bin/deno", "/usr/local/bin/deno"}).
		Sync(ctx)
	if err != nil {
		return err
	}

	// Copy the hello.ts file
	srcDir := client.Host().Directory(".")
	withSrc := deno.WithDirectory("/app", srcDir)

	// Cross-compile for macOS
	macBinary, err := withSrc.WithWorkdir("/app").
		WithExec([]string{"deno", "compile", "--allow-net", "--target", "x86_64-apple-darwin", "hello.ts"}).
		File("/app/hello").
		Sync(ctx)
	if err != nil {
		return err
	}

	// Cross-compile for Windows
	winBinary, err := withSrc.WithWorkdir("/app").
		WithExec([]string{"deno", "compile", "--allow-net", "--target", "x86_64-pc-windows-msvc", "hello.ts"}).
		File("/app/hello.exe").
		Sync(ctx)
	if err != nil {
		return err
	}

	// Compile for Linux
	linuxBinary, err := withSrc.WithWorkdir("/app").
		WithExec([]string{"deno", "compile", "--allow-net", "hello.ts"}).
		File("/app/hello").
		Sync(ctx)
	if err != nil {
		return err
	}

	// Create the final container using distroless
	distroless := client.Container().From("gcr.io/distroless/cc-debian12")
	final := distroless.WithFile("/app/hello", linuxBinary).
		WithWorkdir("/app").
		WithEntrypoint([]string{"/app/hello"})

	// Export the binaries
	mac, err := macBinary.Export(ctx, "hello-mac")
	if err != nil {
		return err
	}
	fmt.Printf("Mac export result %s\n", mac)

	win, err := winBinary.Export(ctx, "hello-win.exe")
	if err != nil {
		return err
	}
	fmt.Printf("Win export result %s\n", win)

	linux, err := linuxBinary.Export(ctx, "hello-linux")
	if err != nil {
		return err
	}
	fmt.Printf("Linux export result %s\n", linux)

	// Export the final container as a tar file
	_, err = final.Export(ctx, "dagger-hello-deno.tar")
	if err != nil {
		return err
	}

	if err := tagDockerImage(); err != nil {
		return err
	}

	fmt.Println("Container image exported as dagger-hello-deno.tar")
	return nil
}
