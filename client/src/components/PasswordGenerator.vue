<template>
  <div class="container-fluid">
    <div class="jumbotron">
      <h1 class="h3 mb-3 font-weight-normal">{{ msg }}</h1>
      <div class="alert alert-info">
        Generate easy to remember (one time) 8 character long passwords. It does
        so by making up passwords with a mix of consonants and vowels pair and
        then appending random digits in the end.
      </div>
      <p class="lead">
        <a
          class="btn btn-primary btn-lg"
          href="https://github.com/goavega-software/passwordgenerator"
          role="button"
          >Github</a
        >
      </p>
    </div>
    <div class="row">
      <div class="col-md-6">
        <form @submit="onSubmit">
          <div class="card">
            <div class="card-header text-center">
              Generate Passwords
            </div>
            <div class="card-body">
              <div class="form-row">
                <div class="form-group col">
                  <label for="numPasswords">Number of passwords</label>
                  <input
                    type="number"
                    class="form-control"
                    v-model.number="input.numPasswords"
                    id="numPasswords"
                    required
                    min="0"
                    max="8"
                  />
                </div>
              </div>
              <div class="custom-control-inline custom-control custom-checkbox">
                <input
                  class="custom-control-input"
                  type="checkbox"
                  id="upperCased"
                  v-model="input.randomUpperCase"
                />
                <label class="custom-control-label" for="upperCased"
                  >Capitalize randomly</label
                >
              </div>
              <div class="custom-control-inline custom-control custom-checkbox">
                <input
                  type="checkbox"
                  class="custom-control-input"
                  v-model="input.leetMode"
                  id="leetMode"
                />
                <label class="custom-control-label" for="leetMode"
                  >1337 Mode</label
                >
              </div>
              <div class="custom-control-inline custom-control custom-checkbox">
                <input
                  class="custom-control-input"
                  type="checkbox"
                  id="symbolize"
                  v-model="input.addSymbol"
                />
                <label class="custom-control-label" for="symbolize"
                  >Suffix symbol</label
                >
              </div>
            </div>
            <div class="card-footer text-muted  text-center">
              <button type="submit" class="btn btn-primary">Generate</button>
            </div>
          </div>
        </form>
      </div>
      <div class="col-md-6">
        <transition
          name="custom-classes-transition"
          enter-active-class="animated tada"
          leave-active-class="animated bounceOutRight"
        >
          <ul class="list-group" v-if="response.passwords.length > 0">
            <li
              v-for="(password, index) in response.passwords"
              v-bind:key="index"
              class="list-group-item"
            >
              <span
                style="display:inline-block"
                :class="{ 'bounce animated': animated[index] }"
                @animationend="animated[index] = false"
                >{{ password }}</span
              >
              <a
                href="javascript:void(0);"
                class="float-right"
                @click="() => copyToClipboard(password, index)"
                >📋</a
              >
            </li>
          </ul>
        </transition>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { GeneratorRequest } from '../models/requests/generator-request';
import { GeneratorService } from '../services/generator-service';
import { PasswordResponse } from '../models/responses/passwords-response';
import { ResponseViewModel } from '../models/responses/response-viewmodel';

@Component
export default class PasswordGenerator extends Vue {
  private input = new GeneratorRequest();
  private response: PasswordResponse = new PasswordResponse();
  private animated: boolean[] = [];
  @Prop() private msg!: string;
  async onSubmit(e: Event) {
    e.stopPropagation();
    e.preventDefault();
    const response = await GeneratorService.generate(this.input);
    // should use Vue.set but laziness bit me
    this.animated = response.data!.passwords.map(item => false);
    this.response = response.data!;
  }

  async copyToClipboard(password: string, index: number) {
    await navigator.clipboard.writeText(password);
    this.animated = this.animated.map((item, i) => i === index);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
ul {
  list-style-type: none;
  padding: 0;
  li {
    display: inline-block;
    margin: 0 10px;
  }
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter,
.fade-leave {
  opacity: 0;
}
</style>
