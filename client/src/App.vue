<template>
  <div id="app">
    <PasswordGenerator msg="Simple Password Generator (TM)" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import PasswordGenerator from './components/PasswordGenerator.vue';
import './assets/scss/root.scss';
import { HttpBaseService } from './services/http-base-service';
@Component({
  components: {
    PasswordGenerator
  }
})
export default class App extends Vue {
  mounted() {
    HttpBaseService.init(
      statusCode => console.log(statusCode),
      () => {}
    );
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #333;
}
</style>
