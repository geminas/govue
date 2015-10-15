<script>

module.exports={
	el:"#show",
	data:{
	target:"{{component}}"
	},
	ready:function(){
	console.log("The {{component}} test main is loaded")
	},
	components:{
	'{{component}}vue':require('./{{component}}.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="{{component}}vue" />
</template>