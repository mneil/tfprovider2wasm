import {client} from "./tfprovider/index.js"

const ec2 = await client("ec2");
const res = ec2.ResourceInstanceCreate(JSON.stringify({
  ami: "ami-2023",
  instance_type: "t3.micro",
  disable_api_termination: true,
}));
console.log(JSON.parse(res));
// console.log(JSON.stringify(JSON.parse(res), undefined, 2))
