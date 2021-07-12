<template>
  <div>
    This is the about page
    <pre>
      {{msg}}
    </pre>
  </div>
</template>


<script>
export default {
  name:'about',
  data (){
    return {
      msg: "This is the message"
    }
  }
}
</script>

<style>

</style>
