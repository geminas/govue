<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				//hello:"world"
			}
		},
		props:[],
		/////////////Life Span/////////////////////
		// created:function(){
		// 	console.log("{{component}} has been created");
		// },
		// beforeCompile:function(){
		// 	console.log("{{component}} beforeCompiled");
		// },
		// compiled:function(){
		// 	console.log("{{component}} has been compiled");
		// },
		// ready:function(){
		// 	console.log("{{component}} has been ready");
		// },
		// attached:function(){
		// 	console.log("{{component}} has been attached");
		// },
		// detached:function(){
		// 	console.log("{{component}} has been detached");
		// },
		// beforeDestory:function(){
		// 	console.log("before {{component}} destoryed");
		// },
		// destoryed:function(){
		// 	console.log("{{component}} has been destoryed");
		// },
		methods:{
			// sayHello:function(){
			// 	console.log("Hello,This is the {{component}} component");
			// }
		},
		computed:{
			// foo:function(){
			// 	return "bar"
			// }
		},
		events:{
			//////////Events 
			// 'hook:created':function(){
			// 	console.log('{{component}} has created again');
			// },
			// greeting:function(msg){
			// 	console.log(msg);
			// },
			// hello:'sayHello'
			///////////Can be Triggered by
			//////vm.$emit('hello')
		},
		watch:{
			//////////Data Watcher
			// 'a':function(val,oldval){
			// 	console.log('new: %s, old: %s', val, oldVal);
			// }
		}

	}
</script>
<template>

<!-- <div>This is component {{component}}</div>
 -->
 </template>
<style type="text/css"></style>