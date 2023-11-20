import {TerraformAwsProvider} from "./wasm_exec.js"
// import * as awsEc2 from "@aws-sdk/client-ec2";

const provider = TerraformAwsProvider();



const services = [
  "ec2",
];

const cached = {}

export async function client(service) {
  if(cached[service]) {
    return cached[service];
  }
  if(!services.includes(service)) {
    throw new Error(`Invalid service name: ${service}`)
  }
  const result = await WebAssembly.instantiateStreaming(fetch(`tfprovider/${service}.wasm`), provider.importObject)
  provider.run(result.instance)
  cached[service] = TerraformAwsProvider[service];
  // TerraformAwsProvider.client[service] = new awsEc2.EC2Client({
  //   region: "us-east-1",
  //   credentials: {
  //     accessKeyId: "",
  //     secretAccessKey: "",
  //     sessionToken: "",
  //   }
  // });
  return cached[service];
}
