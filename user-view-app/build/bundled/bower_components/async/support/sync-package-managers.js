#!/usr/bin/env node
var fs=require("fs"),_=require("lodash"),packageJson=require("../package.json"),IGNORES=["**/.*","node_modules","bower_components","test","tests"],INCLUDES=["lib/async.js","README.md","LICENSE"],REPOSITORY_NAME="caolan/async";packageJson.jam={main:packageJson.main,include:INCLUDES,categories:["Utilities"]},packageJson.spm={main:packageJson.main},packageJson.volo={main:packageJson.main,ignore:IGNORES};var bowerSpecific={moduleType:["amd","globals","node"],ignore:IGNORES,authors:[packageJson.author]},bowerInclude=["name","description","main","keywords","license","homepage","repository","devDependencies"],componentSpecific={repository:REPOSITORY_NAME,scripts:[packageJson.main]},componentInclude=["name","description","version","keywords","license","main"],bowerJson=_.merge({},_.pick(packageJson,bowerInclude),bowerSpecific),componentJson=_.merge({},_.pick(packageJson,componentInclude),componentSpecific);fs.writeFileSync("./bower.json",JSON.stringify(bowerJson,null,2)),fs.writeFileSync("./component.json",JSON.stringify(componentJson,null,2)),fs.writeFileSync("./package.json",JSON.stringify(packageJson,null,2));