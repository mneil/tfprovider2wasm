/**
 * git clones a repository into a directory. Always shallow clones the repo
 * and can take in a list of files or directories to do a sparse checkout when
 * we do not require the entire repository.
 *
 * @param {string} dir - directory to clone the repo into
 * @param {string} repo - owner/name on github
 * @param {string} ref - github ref to clone. Any ref will do. Usually a branch or tag
 * @param {string[]=} files - list of files or directories in the repo to clone. Defaults to everything.
 */
async function git(dir, repo, ref, files = []) {
  console.log(`pulling ${files.join(" ")} from ${repo} into ${dir}`);
  await $`mkdir -p ${dir}`;
  await within(async () => {
    cd(dir);
    await $`git init`;
    if (files.length > 0) {
      await $`git config core.sparsecheckout true`;
      await $`touch .git/info/sparse-checkout`;
      for (const file of files) {
        await $`echo "${file}" >> .git/info/sparse-checkout; rm -rf "${dir}/${file}"; echo "sparse ${file}"`;
      }
    }
    // ignore go test files
    $`echo "!**/*_test.go" >> .git/info/sparse-checkout;`
    await $`git remote add origin https://github.com/${repo}.git`;
    await $`git pull -s theirs --depth 1 origin ${ref}`;
    if (files.length > 0) {
      // only leave .git folder if we cloned the entire repo
      // await $`rm -rf .git`;
    }
  });
}
/**
 * Apply git patches from patchBaseDir directory. Searches patchBaseDir recursively
 * looking for files. All files are treated as git patch files. The subfolder location
 * in patchBaseDir is used as the base directory to apply the patches in. This relative
 * location (toPatch) from the project root should contain a `.git` folder.
 *
 * Do create new patch files cd into the folder to create the patch from and run `git diff > $PATCH_FILE`
 * Move the patch file into the .patches folder with a subdirectory that matches the relative path
 * of the cd directory
 *
 * @param {string} patchBaseDir Base directory to search for patches
 * @param {string} toPatch subfolder to search/apply patches in
 */
async function applyPatches(patchBaseDir, toPatch = "") {
  const dirents = await fs.readdir(path.join(patchBaseDir, toPatch), { withFileTypes: true });
  for (const dirent of dirents) {
    if (dirent.isDirectory()) {
      applyPatches(patchBaseDir, path.join(toPatch, dirent.name));
      continue;
    }
    within(async () => {
      cd(toPatch);
      await $`git apply -p1 ${path.join(patchBaseDir, toPatch, dirent.name)}`;
    });
  }
}
/**
 * Clone the terraform provider aws
 * @param {string} ref
 */
async function cloneTfProviderAws(ref) {
  // clone the terraform aws provider
  await $`rm -rf ./terraform-provider-aws`;
  await git("./terraform-provider-aws", "terraform-providers/terraform-provider-aws", ref, [
    "internal/attrmap",
    "internal/conns",
    "internal/create",
    "internal/enum",
    "internal/envvar",
    "internal/errs",
    "internal/experimental",
    "internal/flex",
    "internal/framework",
    "internal/logging",
    "internal/maps",
    "internal/provider",
    "internal/sdktypes",
    "internal/slices",
    "internal/sweep",
    "internal/tags",
    "internal/tfresource",
    "internal/types",
    "internal/vault",
    "internal/verify",
    "names",
    "version",

    "go.mod",

    // TODO: Don't do this. Just trying some things
    // "internal/service/ec2",
    "internal/service/ec2/arn.go",
    "internal/service/ec2/errors.go",
    "internal/service/ec2/service_package_gen.go",
    "internal/service/ec2/common_schema_data_source.go",
    "internal/service/ec2/filter.go",
    "internal/service/ec2/status.go",
    "internal/service/ec2/consts.go",
    "internal/service/ec2/filters.go",
    // "internal/service/ec2/sweep.go",
    "internal/service/ec2/diff.go",
    "internal/service/ec2/find.go",
    "internal/service/ec2/tags.go",
    "internal/service/ec2/tagsv2_gen.go",
    "internal/service/ec2/ec2_ami.go",
    "internal/service/ec2/ec2_ami_data_source.go",
    // "internal/service/ec2/vpc_security_group_rule_data_source.go",
    // "internal/service/ec2/vpc_security_group_rules_data_source.go",
    "internal/service/ec2/generate.go",
    "internal/service/ec2/tags_gen.go",
    "internal/service/ec2/ec2_instance.go",
    "internal/service/ec2/id.go",
    "internal/service/ec2/validate.go",
    "internal/service/ec2/ec2_instance_migrate.go",
    "internal/service/ec2/list_pages_gen.go",
    "internal/service/ec2/wait.go",
  ]);

  await $`rm -rf ./terraform-provider-aws/conns`
  await $`rm -rf ./terraform-provider-aws/ec2`
  await $`rm -rf ./terraform-provider-aws/provider`

  await $`cp -fr ./misc/awsclient_gen.go ./terraform-provider-aws/internal/conns/awsclient_gen.go`
  await $`cp -fr ./misc/service_packages_gen.go ./terraform-provider-aws/internal/provider/service_packages_gen.go`
  await $`cp -fr ./misc/service_package_gen.go ./terraform-provider-aws/internal/service/ec2/service_package_gen.go`
  await $`cp -fr ./misc/provider.go ./terraform-provider-aws/provider/provider.go`

  await $`mv -f ./terraform-provider-aws/internal/conns ./terraform-provider-aws/conns`
  await $`mv -f ./terraform-provider-aws/internal/service/ec2 ./terraform-provider-aws/ec2`
  await $`mv -f ./terraform-provider-aws/internal/provider ./terraform-provider-aws/provider`

  await $`find ./terraform-provider-aws -type f -exec sed -i 's|github.com/hashicorp/terraform-provider-aws/internal/conns|github.com/hashicorp/terraform-provider-aws/conns|g' {} \\;`
  await $`find ./terraform-provider-aws -type f -exec sed -i 's|github.com/hashicorp/terraform-provider-aws/internal/provider|github.com/hashicorp/terraform-provider-aws/provider|g' {} \\;`



}
/**
 * Clone the terraform plugin sdk so we can patch it
 * @param {string} ref - git ref to clone (tag or branch)
*/
async function cloneTfPluginSdk(ref) {
  // aws helper functions for the go-sdk
  await $`rm -rf ./terraform-plugin-sdk`;
  await git("./terraform-plugin-sdk", "hashicorp/terraform-plugin-sdk", ref);
  // await $`find ./terraform-plugin-sdk -type f -exec sed -i 's|github.com/hashicorp/go-cty/cty|github.com/zclconf/go-cty|g' {} \\;`
  // within(async () => {
  //   cd("./terraform-plugin-sdk")
  //   await $`go get github.com/zclconf/go-cty@v1.14.1`
  // })
}

async function cloneGoogleCty(ref) {
  await $`rm -rf ./go-cty`;
  await git("./go-cty", "zclconf/go-cty", ref);
}
/**
 * Clone the the mock sdk
 * @param {string} ref - git ref to clone (tag or branch)
 */
async function cloneAwsSdkV1(ref) {
  // aws helper functions for the go-sdk
  await $`rm -rf ./aws-sdk-go-mock`;
  await git("./aws-sdk-go-mock", "mneil/aws-sdk-go-mock", ref);
}

async function finalizeClone() {
  // TODO: this isn't working. zx sanitizing it?
  await $`rm -rf mod/**/*_test.go`;
  // this has to happen before we rewrite imports
  await applyPatches(path.resolve(__dirname, "..", "..", ".patches"));
  // await $`mv .terraform-provider-aws/internal/service/** ./terraform-provider-aws/`
  // rewrite imports and patched references / tokens
  // await $`go run main.go -file mod`;
}

async function install() {
  const tfProviderAwsSha = "2d4cce82a8a6a0987f626eb20bf61acfa77e24a0";
  await Promise.all([
    cloneTfProviderAws(tfProviderAwsSha),
    // cloneTfPluginSdk("v2.26.1"),
    // cloneAwsSdkV1("sdkv1"),


    // cloneGoogleCty("7152062cc7333dcdfeed910e7c7f9690276bc2eb"), // v1.14.1
    // cloneAwsSdkV2("4599f78694cabb6853addabc6f92cb197fdb5647")
  ]);


  // await finalizeClone()


  // await $`go mod tidy`
  // await patchWasmExec();
  // TODO: replace this
  // await copyBasicFiles();
  // await $`npm run build`;
}

install()
  .then(() => {
    console.log(chalk.bgGreen("install done"));
  })
  .catch((e) => {
    console.log(chalk.bgRed("install failed"));
    console.error(e);
    process.exit(1);
  });
