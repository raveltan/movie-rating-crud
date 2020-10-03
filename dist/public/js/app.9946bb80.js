(function(e){function t(t){for(var r,n,o=t[0],l=t[1],c=t[2],u=0,v=[];u<o.length;u++)n=o[u],Object.prototype.hasOwnProperty.call(i,n)&&i[n]&&v.push(i[n][0]),i[n]=0;for(r in l)Object.prototype.hasOwnProperty.call(l,r)&&(e[r]=l[r]);d&&d(t);while(v.length)v.shift()();return s.push.apply(s,c||[]),a()}function a(){for(var e,t=0;t<s.length;t++){for(var a=s[t],r=!0,o=1;o<a.length;o++){var l=a[o];0!==i[l]&&(r=!1)}r&&(s.splice(t--,1),e=n(n.s=a[0]))}return e}var r={},i={app:0},s=[];function n(t){if(r[t])return r[t].exports;var a=r[t]={i:t,l:!1,exports:{}};return e[t].call(a.exports,a,a.exports,n),a.l=!0,a.exports}n.m=e,n.c=r,n.d=function(e,t,a){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:a})},n.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var a=Object.create(null);if(n.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)n.d(a,r,function(t){return e[t]}.bind(null,r));return a},n.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],l=o.push.bind(o);o.push=t,o=o.slice();for(var c=0;c<o.length;c++)t(o[c]);var d=l;s.push([0,"chunk-vendors"]),a()})({0:function(e,t,a){e.exports=a("56d7")},"56d7":function(e,t,a){"use strict";a.r(t);a("e260"),a("e6cf"),a("cca6"),a("a79d");var r=a("2b0e"),i=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[e.login?a("Navbar"):e._e(),a("b-loading",{attrs:{"is-full-page":""},model:{value:e.loading,callback:function(t){e.loading=t},expression:"loading"}}),e.login?e._e():a("Login"),e.login?a("router-view"):e._e()],1)},s=[],n=(a("96cf"),a("1da1")),o=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("section",{staticClass:"hero is-light is-fullheight"},[a("div",{staticClass:"hero-body"},[a("div",{staticClass:"container"},[a("div",{staticClass:"columns is-centered"},[a("div",{staticClass:"column is-7-tablet is-6-desktop is-5-widescreen"},[a("form",{staticClass:"box rounded",on:{submit:function(t){return t.preventDefault(),e.login()}}},[a("div",{staticClass:"field has-text-centered title is-3 mt-3"},[e._v(" Movie Rater ")]),a("div",{staticClass:"field has-text-centered title is-6 mb-3"},[e._v(" "+e._s(e.isLogin?"Login to existing account":"Register for an account")+" ")]),e.errorMessage?a("b-notification",{attrs:{type:"is-danger","aria-close-label":"Close notification",role:"alert"}},[e._v(" "+e._s(e.errorMessage)+" ")]):e._e(),a("b-field",{attrs:{label:"Email",type:e.emailError?"is-danger":null,message:e.emailError}},[a("b-input",{attrs:{type:"email",placeholder:"E.g. john@gemail.com",rounded:"","icon-pack":"fas","aria-required":"",icon:"envelope",disabled:e.isLoading,"icon-right":"close-circle","icon-right-clickable":""},on:{"icon-right-click":function(t){e.email=""}},model:{value:e.email,callback:function(t){e.email=t},expression:"email"}})],1),a("b-field",{attrs:{label:"Password",type:e.passwordError?"is-danger":null,message:e.passwordError}},[a("b-input",{attrs:{type:"password",rounded:"",disabled:e.isLoading,placeholder:"********","password-reveal":"","icon-pack":"fas",icon:"key"},model:{value:e.password,callback:function(t){e.password=t},expression:"password"}})],1),e.isLogin?e._e():a("b-field",{attrs:{label:"First Name",type:e.firstNameError?"is-danger":null,message:e.firstNameError}},[a("b-input",{attrs:{type:"text",placeholder:"John",rounded:"","icon-pack":"fas","aria-required":"",disabled:e.isLoading,icon:"user-alt","icon-right":"close-circle","icon-right-clickable":""},on:{"icon-right-click":function(t){e.firstName=""}},model:{value:e.firstName,callback:function(t){e.firstName=t},expression:"firstName"}})],1),e.isLogin?e._e():a("b-field",{attrs:{label:"Last Name",type:e.lastNameError?"is-danger":null,message:e.lastNameError}},[a("b-input",{attrs:{type:"text",placeholder:"Purple",rounded:"","icon-pack":"fas",disabled:e.isLoading,"aria-required":"",icon:"user-alt","icon-right":"close-circle","icon-right-clickable":""},on:{"icon-right-click":function(t){e.lastName=""}},model:{value:e.lastName,callback:function(t){e.lastName=t},expression:"lastName"}})],1),a("div",{staticClass:"field centered"},[a("b-button",{attrs:{rounded:"",expanded:"",outlined:"",disabled:e.isLoading},on:{click:function(t){e.isLogin=!e.isLogin}}},[e._v(" "+e._s(e.isLogin?"Register for an Account":"Already got an account?")+" ")])],1),a("div",{staticClass:"field"},[a("b-button",{attrs:{loading:e.isLoading,"native-type":"submit",rounded:"",expanded:"",type:"is-primary"}},[e._v(" "+e._s(e.isLogin?"Sign In":"Register")+" ")])],1)],1)])])])])])},l=[],c=a("bc3a"),d=a.n(c),u={data:function(){return{isLogin:!1,isLoading:!1,email:"",password:"",emailError:"",passwordError:"",firstName:"",firstNameError:"",lastName:"",lastNameError:"",errorMessage:""}},methods:{login:function(){var e=this;return Object(n["a"])(regeneratorRuntime.mark((function t(){var a,r,i;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.errorMessage="",a=!1,e.password.length<8&&(e.passwordError="Password should be at least 8 characters",a=!0),e.email||(e.emailError="Email should not be empty",a=!0),!e.isLogin||a){t.next=21;break}return e.isLoading=!0,t.prev=6,t.next=9,d.a.post(e.$store.state.baseUrl+"/api/login",{email:e.email,password:e.password});case 9:r=t.sent,e.$store.commit("login",{token:r.data.token,refresh:r.data.refresh}),t.next=16;break;case 13:t.prev=13,t.t0=t["catch"](6),t.t0&&(e.errorMessage="Unable to login, please check your credentials");case 16:return t.prev=16,e.isLoading=!1,t.finish(16);case 19:t.next=38;break;case 21:if(e.firstName.length<3&&(e.firstNameError="First Name should be at least 3 characters",a=!0),e.lastName.length<3&&(e.lastNameError="Last Name should be at least 3 characters",a=!0),a){t.next=38;break}return e.isLoading=!0,t.prev=25,t.next=28,d.a.post(e.$store.state.baseUrl+"/api/register",{email:e.email,password:e.password,firstName:e.firstName,lastName:e.lastName});case 28:i=t.sent,e.$store.commit("login",{token:i.data.token,refresh:i.data.refresh}),t.next=35;break;case 32:t.prev=32,t.t1=t["catch"](25),t.t1&&("User exists"==t.t1.response.data.error?e.errorMessage="User with with this email is already registered.":e.errorMessage="Unable to Register, please check your credentials");case 35:return t.prev=35,e.isLoading=!1,t.finish(35);case 38:case"end":return t.stop()}}),t,null,[[6,13,16,19],[25,32,35,38]])})))()}}},v=u,m=a("2877"),p=Object(m["a"])(v,o,l,!1,null,null,null),f=p.exports,g=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("b-navbar",{attrs:{shadow:"","fixed-top":""}},[a("template",{slot:"brand"},[a("b-navbar-item",{attrs:{tag:"router-link",to:{path:"/"}}},[a("h1",{staticClass:"title is-4"},[e._v("Movie Rater")])])],1),a("template",{slot:"end"},[a("b-navbar-item",{attrs:{tag:"div"}},[a("div",{staticClass:"buttons"},[a("b-button",{attrs:{rounded:"","icon-pack":"fas","icon-right":"sign-out-alt",expanded:""},on:{click:function(t){return e.$store.commit("logout")}}},[a("strong",[e._v("Sign Out")])])],1)])],1)],2)},b=[],h={},w=h,_=Object(m["a"])(w,g,b,!1,null,null,null),k=_.exports,C={components:{Login:f,Navbar:k},data:function(){return{login:!1,loading:!0}},created:function(){this.initLogin()},methods:{initLogin:function(){var e=this;return Object(n["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,e.$store.dispatch("getLogin");case 2:e.login=null!=e.$store.state.token,e.$store.state.loading=!1;case 4:case"end":return t.stop()}}),t)})))()}},watch:{"$store.state.token":function(){this.$store.state.token&&this.$store.state.refresh?this.login=!0:this.login=!1},"$store.state.loading":function(){this.loading=this.$store.state.loading}}},y=C,x=Object(m["a"])(y,i,s,!1,null,null,null),R=x.exports,N=a("8c4f"),$=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("section",{staticClass:"section"},[a("div",{staticClass:"container"},[a("div",{staticClass:"level"},[a("div",{staticClass:"left"},[a("div",{staticClass:"item"},[a("h1",{staticClass:"title is-5 mb-3"},[e._v(" Movie ("+e._s(e.movie.length)+") ")])])]),a("div",{staticClass:"right"},[a("div",{staticClass:"item is-grouped is-grouped-centered"},[a("b-button",{staticClass:"mr-2",attrs:{outlined:"",type:"is-primary is-light"},on:{click:function(t){e.addDialog=!0}}},[e._v(" Add Movie ")]),a("b-button",{attrs:{outlined:""},on:{click:function(t){return e.initData()}}},[e._v(" Refresh ")])],1)])]),a("b-modal",{attrs:{"has-modal-card":"","trap-focus":"","destroy-on-hide":!1,"aria-role":"dialog","aria-modal":""},model:{value:e.addDialog,callback:function(t){e.addDialog=t},expression:"addDialog"}},[[a("form",{attrs:{action:""},on:{submit:function(t){return t.preventDefault(),e.addMovie(t)}}},[a("div",{staticClass:"modal-card",staticStyle:{width:"auto"}},[a("header",{staticClass:"modal-card-head"},[a("p",{staticClass:"modal-card-title"},[e._v("Add New Movie")]),a("button",{staticClass:"delete",attrs:{type:"button"},on:{click:function(t){e.addDialog=!1}}})]),a("section",{staticClass:"modal-card-body"},[a("b-field",{attrs:{label:"Movie Name"}},[a("b-input",{attrs:{type:"text",placeholder:"Hua Mu Lan"},model:{value:e.newMovie,callback:function(t){e.newMovie=t},expression:"newMovie"}})],1)],1),a("footer",{staticClass:"modal-card-foot"},[a("button",{staticClass:"button",attrs:{type:"button"},on:{click:function(t){e.addDialog=!1}}},[e._v(" Close ")]),a("b-button",{attrs:{type:"is-primary","native-type":"submit"}},[e._v("Add")])],1)])])]],2),e.error?a("b-notification",{attrs:{type:"is-danger","aria-close-label":"Close notification",role:"alert"}},[e._v(e._s(e.error)+" ")]):e._e(),a("b-table",{attrs:{data:e.movie,"sticky-header":"",striped:"",height:"50vh","default-sort":["ID","asc"],paginated:"","icon-pack":"fas"}},[a("b-table-column",{attrs:{field:"ID",label:"ID",sortable:"",width:"40",centered:"",numeric:""},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.ID)+" ")]}}])}),a("b-table-column",{attrs:{field:"Name",label:"Movie Name",width:"600"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.Name)+" ")]}}])}),a("b-table-column",{attrs:{field:"Rating",label:"Rating",sortable:"",centered:"",width:"40",numeric:""},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",{staticClass:"tag is-info is-light"},[a("strong",{staticClass:"is-white"},[e._v(e._s(t.row.Rating))]),a("span",{staticClass:"icon "},[a("i",{staticClass:"fas fa-star"})])])]}}])}),a("b-table-column",{attrs:{label:"Action",centered:"",width:"20"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("b-button",{attrs:{size:"is-small",rounded:""},on:{click:function(a){return e.$router.push("/review/"+t.row.ID)}}},[e._v(" Review ")])]}}])})],1)],1)])},L=[],D={data:function(){return{movie:[],error:"",addDialog:!1,newMovie:"",unsubs:null}},created:function(){this.initData()},methods:{addMovie:function(){var e=this;return Object(n["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.addDialog=!1,""!=e.newMovie){t.next=3;break}return t.abrupt("return");case 3:return e.$store.state.loading=!0,t.prev=4,t.next=7,d.a.post(e.$store.state.baseUrl+"/api/movie/add",{name:e.newMovie});case 7:e.initData(),t.next=13;break;case 10:t.prev=10,t.t0=t["catch"](4),t.t0&&(e.error="Problem communication with the server");case 13:return t.prev=13,e.newMovie="",t.finish(13);case 16:case"end":return t.stop()}}),t,null,[[4,10,13,16]])})))()},initData:function(){var e=this;return Object(n["a"])(regeneratorRuntime.mark((function t(){var a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return d.a.defaults.headers.common={Authorization:"Bearer ".concat(e.$store.state.token)},e.$store.state.loading=!0,t.prev=2,t.next=5,d.a.get(e.$store.state.baseUrl+"/api/movies");case 5:a=t.sent,a&&a.data&&(a.data.length>0?e.movie=a.data:e.movie=[]),t.next=13;break;case 9:t.prev=9,t.t0=t["catch"](2),console.log({e:t.t0}),t.t0&&(e.error="Problem communication with the server");case 13:return t.prev=13,e.$store.state.loading=!1,t.finish(13);case 16:case"end":return t.stop()}}),t,null,[[2,9,13,16]])})))()}}},M=D,E=Object(m["a"])(M,$,L,!1,null,null,null),O=E.exports,S=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("div",{staticClass:"modal"},[a("div",{staticClass:"modal-background"}),a("div",{staticClass:"modal-card"},[a("header",{staticClass:"modal-card-head"},[a("p",{staticClass:"modal-card-title"},[e._v("Add new Review")]),a("button",{staticClass:"delete",attrs:{"aria-label":"close"},on:{click:function(t){return e.hideAddDialog()}}})]),a("section",{staticClass:"modal-card-body"},[a("form",{on:{submit:function(t){return t.preventDefault(),e.addReview(t)}}},[a("div",{staticClass:"field"},[a("label",{staticClass:"label",attrs:{for:"name"}},[e._v("Review")]),a("div",{staticClass:"control"},[a("input",{directives:[{name:"model",rawName:"v-model",value:e.newReview,expression:"newReview"}],staticClass:"input",attrs:{type:"text",name:"name",placeholder:"Very Good",required:""},domProps:{value:e.newReview},on:{input:function(t){t.target.composing||(e.newReview=t.target.value)}}})])]),a("div",{staticClass:"field"},[a("label",{staticClass:"label",attrs:{for:"name"}},[e._v("Rating")]),a("div",{staticClass:"control"},[a("input",{directives:[{name:"model",rawName:"v-model",value:e.newRating,expression:"newRating"}],staticClass:"input",attrs:{type:"number",name:"nrating",placeholder:"5",required:"",min:"0",max:"5"},domProps:{value:e.newRating},on:{input:function(t){t.target.composing||(e.newRating=t.target.value)}}})])]),a("div",{staticClass:"field is-grouped"},[e._m(0),a("div",{staticClass:"control"},[a("button",{staticClass:"button is-light",attrs:{type:"button"},on:{click:function(t){return e.hideAddDialog()}}},[e._v(" Cancel ")])])])])])])]),a("section",{staticClass:"section"},[a("div",{staticClass:"container"},[e.error?a("b-notification",{attrs:{type:"is-danger","aria-close-label":"Close notification",role:"alert"}},[e._v(e._s(e.error)+" ")]):e._e(),a("div",{staticClass:"level"},[e.notFound?e._e():a("div",{staticClass:"level-left"},[a("div",{staticClass:"level-item"},[a("h1",{staticClass:"title"},[e._v(e._s(e.reviews.name))])]),a("div",{staticClass:"level-item"},[a("p",{staticClass:"subtitle"},[e._v(e._s(e.reviews.rating)+" stars")])])]),e.notFound?a("div",{staticClass:"level-left"},[e._m(1)]):e._e(),a("div",{staticClass:"level-right"},[a("div",{staticClass:"level-item"},[a("button",{staticClass:"button",on:{click:function(t){return e.showAddDialog()}}},[e._v(" Add Review ")])])])]),e.notFound?e._e():a("div",{staticClass:"columns"},[a("div",{staticClass:"column"},[a("div",{staticClass:"card is-6-tablet is-5-desktop is-3-widescreen"},e._l(e.reviews.review,(function(t){return a("div",{key:t.Name+t.Review+t.Rating,staticClass:"card-content"},[a("p",{staticClass:"title"},[e._v("“"+e._s(t.Review)+"”")]),a("p",{staticClass:"subtitle"},[e._v(" "+e._s(t.Name)+" ("+e._s(t.Rating)+" stars) ")])])})),0)])])],1)])])},j=[function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"control"},[a("button",{staticClass:"button is-success",attrs:{type:"submit",name:"movie"}},[e._v(" Add Review ")])])},function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"level-item"},[a("p",{staticClass:"title is-4"},[e._v("Be the first to add review")])])}],A={data:function(){return{newReview:"",newRating:5,error:"",reviews:{},notFound:!1}},created:function(){this.initData()},methods:{showAddDialog:function(){var e=document.querySelector(".modal");e.classList.add("is-active")},hideAddDialog:function(){var e=document.querySelector(".modal");e.classList.remove("is-active")},addReview:function(){var e=this;return Object(n["a"])(regeneratorRuntime.mark((function t(){var a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.hideAddDialog(),!(e.newReview.length<1)||e.newRating){t.next=3;break}return t.abrupt("return");case 3:return e.$store.state.loading=!0,t.prev=4,a={review:e.newReview,rating:parseInt(e.newRating)},t.next=8,d.a.post(e.$store.state.baseUrl+"/api/review/"+e.$route.params.id+"/add",a);case 8:e.error="",e.newReview="",e.newRating=5,e.initData(),t.next=17;break;case 14:t.prev=14,t.t0=t["catch"](4),t.t0&&(e.error="Problem communication with the server");case 17:case"end":return t.stop()}}),t,null,[[4,14]])})))()},initData:function(){var e=this;return Object(n["a"])(regeneratorRuntime.mark((function t(){var a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return d.a.defaults.headers.common={Authorization:"Bearer ".concat(e.$store.state.token)},e.$store.state.loading=!0,t.prev=2,t.next=5,d.a.get(e.$store.state.baseUrl+"/api/review/"+e.$route.params.id);case 5:a=t.sent,a&&(e.reviews=a.data),e.error="",e.notFound=!1,t.next=19;break;case 11:if(t.prev=11,t.t0=t["catch"](2),"404"!=t.t0.response.status){t.next=18;break}return e.notFound=!0,t.abrupt("return");case 18:e.error="Problem communication with the server";case 19:return t.prev=19,e.$store.state.loading=!1,t.finish(19);case 22:case"end":return t.stop()}}),t,null,[[2,11,19,22]])})))()}}},P=A,I=Object(m["a"])(P,S,j,!1,null,null,null),U=I.exports;r["a"].use(N["a"]);var F=[{path:"/",name:"Home",component:O},{path:"/review/:id",component:U,name:"Review"}],q=new N["a"]({routes:F}),z=q,B=a("2f62"),J=a("289d");a("5abe");r["a"].use(J["a"]),r["a"].config.productionTip=!1,r["a"].use(B["a"]);var T=new B["a"].Store({state:{token:null,refresh:null,baseUrl:"http://localhost:3000",loading:!0},mutations:{login:function(e,t){localStorage.setItem("token",t.token),localStorage.setItem("refresh",t.refresh),e.token=t.token,e.refresh=t.refresh},logout:function(e){localStorage.removeItem("token"),localStorage.removeItem("refresh"),e.refresh=null,e.token=null}},actions:{getLogin:function(e){localStorage.getItem("token")&&localStorage.getItem("refresh")&&e.commit("login",{token:localStorage.getItem("token"),refresh:localStorage.getItem("refresh")})}}});new r["a"]({router:z,render:function(e){return e(R)},store:T}).$mount("#app")}});
//# sourceMappingURL=app.9946bb80.js.map