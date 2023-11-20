async function build() {
  // compile
  await $`mkdir -p dist`
  await within(async () => {
    cd("./tfprovider2wasm")
    await $`GOOS=js GOARCH=wasm go build -o ec2.wasm -v`
  })
  await $`cp -f tfprovider2wasm/js/index.html dist`
  await $`cp -f tfprovider2wasm/js/index.js dist`
  await $`cp -rf tfprovider2wasm/js/tfprovider dist/tfprovider`
  await $`cp -f tfprovider2wasm/ec2.wasm dist/tfprovider`
}

build()
  .then(() => {
    console.log(chalk.bgGreen("build done"));
  })
  .catch((e) => {
    console.log(chalk.bgRed("build failed"));
    console.error(e);
    process.exit(1);
  });
