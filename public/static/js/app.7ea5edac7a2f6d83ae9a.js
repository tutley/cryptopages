webpackJsonp([1],{"/626":function(e,t){},"2V2b":function(e,t,s){"use strict";var n=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("div",[e._v("This is the About Page")])])},r=[],a={render:n,staticRenderFns:r};t.a=a},"3TcT":function(e,t){},"6nfk":function(e,t){},"8VSw":function(e,t){},"8dY9":function(e,t){},"90VL":function(e,t){},"9DHQ":function(e,t){},BCfr:function(e,t){},FbEL:function(e,t){},"HNa/":function(e,t){},Ij1S:function(e,t){},JLzs:function(e,t){},M8Wo:function(e,t){},"N+IJ":function(e,t){},NHnr:function(e,t,s){"use strict";function n(e){s("8dY9")}function r(e){s("Ij1S")}function a(e){s("M8Wo")}Object.defineProperty(t,"__esModule",{value:!0});var i=(s("s6ZO"),s("7+uW")),o=s("M+UZ"),l=s("IcMm"),u=s("rPQa"),c=s("fYhH"),v=s("7Q1V"),d=s("7M7f"),m=s("f/u0"),h=s("z8aH"),f=s("+9ps"),p=s("TWha"),b=s("XRgG"),g=s("DDBM"),_=s("QLUw"),k=s("pqui"),w=s("+1ch"),x=s("7gKz"),y=s("/yi5"),C=s("BOXn"),$=s("JUsQ"),L=s("MPXs"),R=s("6VHA"),D={name:"app",data:function(){return{drawer:!1}},computed:{userIsAuthenticated:function(){return this.$store.getters.isLoggedIn},menuItems:function(){var e=[{icon:"home",title:"Home",link:"/"},{icon:"info",title:"About",link:"/about"},{icon:"lock_open",title:"Sign in",link:"/signin"},{icon:"add_circle",title:"Sign Up",link:"/signup"}];return this.userIsAuthenticated&&(e=[{icon:"home",title:"Home",link:"/"},{icon:"info",title:"About",link:"/about"},{icon:"search",title:"Search Listings",link:"/search"}]),e}},methods:{jump:function(e){this.$router.push(e)},onLogout:function(){this.$store.dispatch("logout"),this.$router.push("/")}},beforeMount:function(){var e=document.getElementById("pwaloader");e&&e.remove()}},j=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-app",{attrs:{id:"cryptopages"}},[s("v-navigation-drawer",{attrs:{app:"",fixed:"",light:""},model:{value:e.drawer,callback:function(t){e.drawer=t},expression:"drawer"}},[s("v-list",[e._l(e.menuItems,function(t){return s("v-list-tile",{key:t.title,attrs:{exact:"",to:t.link}},[s("v-list-tile-action",[s("v-icon",[e._v(e._s(t.icon))])],1),e._v(" "),s("v-list-tile-content",[e._v(e._s(t.title))])],1)}),e._v(" "),e.userIsAuthenticated?s("v-list-tile",{on:{click:e.onLogout}},[s("v-list-tile-action",[s("v-icon",[e._v("exit_to_app")])],1),e._v(" "),s("v-list-tile-content",[e._v("Logout")])],1):e._e()],2)],1),e._v(" "),s("v-toolbar",{attrs:{app:"",dark:"",color:"primary"}},[s("v-toolbar-side-icon",{on:{click:function(t){t.stopPropagation(),e.drawer=!e.drawer}}}),e._v(" "),s("v-toolbar-title",[e._v("\n        Cryptopages\n      ")]),e._v(" "),s("v-spacer"),e._v(" "),s("v-btn",{attrs:{flat:"",to:"/about"}},[e._v("About")])],1),e._v(" "),s("v-content",{ref:"content"},[s("v-fab-transition",[s("v-btn",{attrs:{fab:"",color:"accent",bottom:"",right:"",fixed:"",small:""},on:{click:function(t){e.jump("/search")}}},[s("v-icon",{attrs:{dark:""}},[e._v("search")])],1)],1),e._v(" "),s("div",[e._v(" ")]),e._v(" "),s("router-view"),e._v(" "),s("v-footer",{staticClass:"pa-3"},[s("div",[s("a",{attrs:{href:"https://tomutley.com/"}},[e._v("Tom Utley")])]),e._v(" "),s("v-spacer")],1)],1)],1)},I=[],V={render:j,staticRenderFns:I},U=V,P=s("VU/8"),S=n,E=P(D,U,!1,S,null,null),A=E.exports,N=s("/ocq"),F={name:"hello",data:function(){return{msg:"Its like the Yellow Pages but for Crypto!"}}},T=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",[s("v-flex",[s("div",{staticClass:"hello"},[s("p",{staticClass:"subheading"},[e._v("\n            Hello, welcome to cryptopages.\n            ")]),e._v(" "),s("p",[e._v("\n              This site is not ready to use yet. When it's ready, we'll update this page.\n            ")])])])],1)],1)},H=[],W={render:T,staticRenderFns:H},q=W,B=s("VU/8"),J=r,M=B(F,q,!1,J,null,null),O=M.exports,z=s("Dd8w"),Y=s.n(z),Q=s("mtWM"),X=s.n(Q),Z=null!=localStorage.getItem("token"),G="";Z&&(G="Bearer "+localStorage.getItem("token"));var K={baseURL:"https://cryptopages.club/api/",headers:Y()({},Z?{Authorization:G}:{})},ee=X.a.create(K),te={methods:{displayDetails:function(e){this.$router.push({name:"User Detail",params:{username:e}})}},data:function(){return{users:[],errors:[],loading:!1}},created:function(){var e=this;this.loading=!0,ee.get("user").then(function(t){e.users=t.data,e.loading=!1}).catch(function(t){e.errors.push(t),e.loading=!1})}},se=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm6:"","offset-sm3":""}},[s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.loading,expression:"loading"}],attrs:{indeterminate:!0}}),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.users.length<1&&!e.loading,expression:"users.length < 1 && !loading"}],attrs:{color:"info",icon:"info",value:"true"}},[e._v("\n      Sorry, no matching users.\n      ")]),e._v(" "),s("v-card",{directives:[{name:"show",rawName:"v-show",value:e.errors.length<1,expression:"errors.length < 1"}]},[s("v-list",{attrs:{"three-line":""}},[s("v-subheader",[e._v("Results:")]),e._v(" "),e._l(e.users,function(t){return[s("v-divider"),e._v(" "),s("v-list-tile",{key:t.username,on:{click:function(s){e.displayDetails(t.username)}}},[s("v-list-tile-content",[s("v-list-tile-title",{domProps:{innerHTML:e._s(t.username)}}),e._v(" "),s("v-list-tile-sub-title",[e._v(e._s(t.jobCategory)+" - "+e._s(t.skills.join(" ")))])],1)],1)]})],2)],1),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.errors.length>0,expression:"errors.length > 0"}],attrs:{color:"error",icon:"warning",value:"true"}},e._l(e.errors,function(t,n){return s("p",{key:n},[e._v("\n          "+e._s(t.message)+"\n        ")])}))],1)],1)],1)},ne=[],re={render:se,staticRenderFns:ne},ae=re,ie=s("VU/8"),oe=ie(te,ae,!1,null,null,null),le=oe.exports,ue=s("woOf"),ce=s.n(ue),ve={data:function(){return{user:{},nameRules:[],emailRules:[function(e){return 0===e.length||/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(e)||"E-mail must be valid"}],locationRules:[],skillsRules:[],jobDescriptionRules:[],otherCoinRules:[],errors:[],valid:!0,editable:!1,showDelete:!1,_originaluser:{},loading:!1,processing:!1}},computed:{thisUser:function(){return this.user.username===this.loggedInUser.username},loggedInUser:function(){return this.$store.getters.user},coinLabels:function(){return this.$store.getters.coinLabels},jobCategories:function(){return this.$store.getters.jobCategories}},methods:{doEdit:function(e){this._originaluser=ce()({},e),this.editable=!0},cancelEdit:function(){ce()(this.user,this._originaluser),this.editable=!1},saveUpdate:function(){var e=this;this.$refs.form.validate()&&(this.processing=!0,ee.put("user/"+this.$route.params.username,this.user).then(function(t){e.editable=!1,e.processing=!1}).catch(function(t){e.errors.push(t),e.processing=!1}))},doDelete:function(){this.showDelete=!0},cancelDelete:function(){this.showDelete=!1},sendDelete:function(){var e=this;this.processing=!0,ee.delete("user/"+this.$route.params.username).then(function(t){e.processing=!1,e.$router.push({name:"Home"})}).catch(function(t){e.errors.push(t),e.processing=!1})}},beforeCreate:function(){var e=this;this.loading=!0,ee.get("user/"+this.$route.params.username).then(function(t){e.user=t.data,e.loading=!1}).catch(function(t){e.errors.push(t),e.loading=!1})}},de=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm8:"","offset-sm2":""}},[s("v-card",[s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.loading,expression:"loading"}],attrs:{indeterminate:!0}}),e._v(" "),s("v-card-title",{attrs:{headline:""}},[s("h2",[e._v(e._s(e.user.username))])]),e._v(" "),e.editable?e._e():s("v-list",{attrs:{"two-line":""}},[s("v-list-tile",[e.user.name?s("v-list-tile-content",[s("v-list-tile-sub-title",[e._v("Name")]),e._v(" "),s("v-list-tile-title",[e._v(e._s(e.user.name))])],1):e._e()],1),e._v(" "),e.user.email?s("v-list-tile",[e.user.email.makePublic?s("v-list-tile-content",[s("v-list-tile-sub-title",[e._v("Email")]),e._v(" "),s("v-list-tile-title",[e._v(e._s(e.user.email.value))])],1):e._e()],1):e._e(),e._v(" "),e.user.location?s("v-list-tile",[e.user.location.makePublic?s("v-list-tile-content",[s("v-list-tile-sub-title",[e._v("Location")]),e._v(" "),s("v-list-tile-title",[e._v(e._s(e.user.location.value))])],1):e._e()],1):e._e(),e._v(" "),s("v-list-tile",[s("v-list-tile-content",[s("v-list-tile-sub-title",[e._v("Job Category")]),e._v(" "),s("v-list-tile-title",[e._v(e._s(e.jobCategories.find(function(t){return t.val===e.user.jobCategory}).name))])],1)],1),e._v(" "),s("v-list-tile",[e.user.jobDescription?s("v-list-tile-content",[s("v-list-tile-sub-title",[e._v("Job Description")]),e._v(" "),s("v-list-tile-title",[e._v(e._s(e.user.jobDescription))])],1):e._e()],1),e._v(" "),s("v-list-tile",[e.user.skills.length>0?s("v-list-tile-content",[s("v-list-tile-sub-title",[e._v("Skills")]),e._v(" "),s("v-list-tile-title",[e._v("\n                "+e._s(e.user.name)+"\n\n              ")])],1):e._e()],1),e._v(" "),s("v-list-tile",[s("v-list-tile-content",[s("v-list-tile-sub-title",[e._v("Coins Accepted")]),e._v(" "),s("v-list-tile-title",[e._v("\n                "+e._s(e.user.name)+"\n              ")])],1)],1)],1),e._v(" "),e.editable?s("v-card-text",[s("v-form",{ref:"form",attrs:{"lazy-validation":""},model:{value:e.valid,callback:function(t){e.valid=t},expression:"valid"}},[s("v-text-field",{attrs:{label:"Name",rules:e.nameRules},model:{value:e.user.name,callback:function(t){e.$set(e.user,"name",t)},expression:"user.name"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Available? (included in search results)"},model:{value:e.user.available,callback:function(t){e.$set(e.user,"available",t)},expression:"user.available"}}),e._v(" "),s("v-text-field",{attrs:{label:"Email",rules:e.emailRules,"validate-on-blur":"",hint:"Let others know the best email to reach you with. You have the option to hide this."},model:{value:e.user.email.value,callback:function(t){e.$set(e.user.email,"value",t)},expression:"user.email.value"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Display Email?"},model:{value:e.user.email.makePublic,callback:function(t){e.$set(e.user.email,"makePublic",t)},expression:"user.email.makePublic"}}),e._v(" "),s("v-text-field",{attrs:{label:"location",rules:e.locationRules,"validate-on-blur":"",hint:"Let others know where you are located. You have the option to hide this."},model:{value:e.user.location.value,callback:function(t){e.$set(e.user.location,"value",t)},expression:"user.location.value"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Display Location?"},model:{value:e.user.location.makePublic,callback:function(t){e.$set(e.user.location,"makePublic",t)},expression:"user.location.makePublic"}}),e._v(" "),s("v-select",{attrs:{items:e.jobCategories,label:"Job Category","single-line":"","item-text":"name","item-value":"val",hint:"What type of work are you interested in?"},model:{value:e.user.jobCategory,callback:function(t){e.$set(e.user,"jobCategory",t)},expression:"user.jobCategory"}}),e._v(" "),s("v-text-field",{attrs:{label:"Job Description",rules:e.jobDescriptionRules,hint:"What type of work are you looking for? Full-time, part-time, contract, big projects, small projects, etc"},model:{value:e.user.jobDescription,callback:function(t){e.$set(e.user,"jobDescription",t)},expression:"user.jobDescription"}}),e._v(" "),s("v-text-field",{attrs:{label:"Skills",rules:e.skillsRules,hint:"What skills do you bring to the table? Use spaces to separate skills."},model:{value:e.user.skills,callback:function(t){e.$set(e.user,"skills",t)},expression:"user.skills"}}),e._v(" "),s("span",{staticClass:"subheading"},[e._v("Coins Accepted")]),s("br"),e._v(" "),s("v-layout",{attrs:{row:"",wrap:""}},e._l(e.user.coins,function(t,n){return s("v-flex",{key:n,staticClass:"xs4 sm3"},[s("v-checkbox",{attrs:{label:e.coinLabels[n]},model:{value:e.user.coins[n],callback:function(t){e.$set(e.user.coins,n,t)},expression:"user.coins[key]"}})],1)})),e._v(" "),s("v-text-field",{directives:[{name:"show",rawName:"v-show",value:e.user.coins.other,expression:"user.coins.other"}],attrs:{label:"Other Coin Name",rules:e.otherCoinRules,hint:"What is the name of the other crypto currency you accept?"},model:{value:e.user.otherCoin,callback:function(t){e.$set(e.user,"otherCoin",t)},expression:"user.otherCoin"}})],1)],1):e._e(),e._v(" "),s("v-card-actions",{directives:[{name:"show",rawName:"v-show",value:!e.editable&&e.thisUser,expression:"!editable && thisUser"}]},[s("v-btn",{attrs:{small:"",color:"secondary"},on:{click:function(t){e.doEdit(e.user)}}},[e._v("EDIT")]),e._v(" "),s("v-spacer"),e._v(" "),s("v-btn",{attrs:{small:"",color:"error"},on:{click:e.doDelete}},[e._v("DELETE")])],1),e._v(" "),s("v-card-actions",{directives:[{name:"show",rawName:"v-show",value:e.editable,expression:"editable"}]},[s("v-btn",{attrs:{small:"",color:"success"},on:{click:e.saveUpdate}},[e._v("SAVE")]),e._v(" "),s("v-spacer"),e._v(" "),s("v-btn",{attrs:{small:"",color:"warning"},on:{click:e.cancelEdit}},[e._v("CANCEL")])],1),e._v(" "),s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.processing,expression:"processing"}],attrs:{indeterminate:!0}})],1),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.errors.length>0,expression:"errors.length > 0"}],attrs:{color:"error",icon:"warning",value:"true"}},e._l(e.errors,function(t,n){return s("p",{key:n},[e._v("\n          "+e._s(t.message)+"\n        ")])}))],1)],1),e._v(" "),s("v-dialog",{model:{value:e.showDelete,callback:function(t){e.showDelete=t},expression:"showDelete"}},[s("v-card",[s("v-card-title",{staticClass:"headline"},[e._v("Really delete user?")]),e._v(" "),s("v-card-text",[e._v("Are you sure? Think of the children!")]),e._v(" "),s("v-card-actions",[s("v-spacer"),e._v(" "),s("v-btn",{attrs:{color:"primary",flat:""},on:{click:e.cancelDelete}},[e._v("CANCEL")]),e._v(" "),s("v-btn",{attrs:{color:"error",flat:""},on:{click:e.sendDelete}},[e._v("DELETE")])],1)],1)],1)],1)},me=[],he={render:de,staticRenderFns:me},fe=he,pe=s("VU/8"),be=pe(ve,fe,!1,null,null,null),ge=be.exports,_e={data:function(){var e=this;return{sending:!1,usernameAvailable:!0,user:{username:"",password:"",available:!0,name:"",email:{value:"",makePublic:!1},location:{value:"",makePublic:!1},skills:"",jobCategory:"others",jobDescription:"",coins:{bcc:!1,btc:!1,eth:!1,ltc:!1,neo:!1,other:!1,xlm:!1,xrp:!1},otherCoin:""},usernameRules:[function(e){return!!e||"Username is required!"},function(e){return e&&e.length<=60||"Username must be less than 60 characters"},function(t){return e.usernameAvailable||"Sorry, that username is taken"}],passwordRules:[function(e){return!!e||"Password is required"}],nameRules:[],emailRules:[function(e){return 0===e.length||/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(e)||"E-mail must be valid"}],skillsRules:[],locationRules:[],jobDescriptionRules:[],otherCoinRules:[],errors:[],valid:!0}},computed:{coinLabels:function(){return this.$store.getters.coinLabels},jobCategories:function(){return this.$store.getters.jobCategories}},methods:{checkUsername:function(e){var t=this;ee.get("user/checkUsername/"+this.user.username).then(function(e){t.usernameAvailable=!1,t.$refs.username.validate()}).catch(function(e){t.usernameAvailable=!0,t.$refs.username.validate()})},submit:function(){var e=this,t=this.user;t.skills=t.skills.split(" "),this.$refs.form.validate()&&(this.sending=!0,ee.post("/user",t).then(function(t){e.sending=!1,e.$router.push({name:"Signin"})}).catch(function(t){e.sending=!1,e.user.skills="",e.errors.push(t)}))},clear:function(){this.$refs.form.reset()}}},ke=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm8:"","offset-sm2":""}},[s("p",{staticClass:"display-2"},[e._v("Sign Up!")]),e._v(" "),s("p",{staticClass:"body-1"},[e._v('Provide as much or as little information as you wish. You must create \n        an account to be able to search and view listings, but you don\'t have to make yourself \n        visible to others. Simply check "available" if you want others to be able to find you!')]),e._v(" "),s("v-form",{ref:"form",attrs:{"lazy-validation":""},model:{value:e.valid,callback:function(t){e.valid=t},expression:"valid"}},[s("v-text-field",{ref:"username",attrs:{label:"Username",autofocus:"",rules:e.usernameRules,counter:60,hint:"This is your identifier on this site. It must be unique, and others will see it.",required:""},on:{blur:e.checkUsername},model:{value:e.user.username,callback:function(t){e.$set(e.user,"username",t)},expression:"user.username"}}),e._v(" "),s("v-text-field",{attrs:{label:"Password",rules:e.passwordRules,required:"",type:"password"},model:{value:e.user.password,callback:function(t){e.$set(e.user,"password",t)},expression:"user.password"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Available? (included in search results)"},model:{value:e.user.available,callback:function(t){e.$set(e.user,"available",t)},expression:"user.available"}}),e._v(" "),s("v-text-field",{attrs:{label:"Name",rules:e.nameRules,hint:"Put your name here."},model:{value:e.user.name,callback:function(t){e.$set(e.user,"name",t)},expression:"user.name"}}),e._v(" "),s("v-text-field",{attrs:{label:"Email",rules:e.emailRules,"validate-on-blur":"",hint:"Let others know the best email to reach you with. You have the option to hide this."},model:{value:e.user.email.value,callback:function(t){e.$set(e.user.email,"value",t)},expression:"user.email.value"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Display Email?"},model:{value:e.user.email.makePublic,callback:function(t){e.$set(e.user.email,"makePublic",t)},expression:"user.email.makePublic"}}),e._v(" "),s("v-text-field",{attrs:{label:"location",rules:e.locationRules,"validate-on-blur":"",hint:"Let others know where you are located. You have the option to hide this."},model:{value:e.user.location.value,callback:function(t){e.$set(e.user.location,"value",t)},expression:"user.location.value"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Display Location?"},model:{value:e.user.location.makePublic,callback:function(t){e.$set(e.user.location,"makePublic",t)},expression:"user.location.makePublic"}}),e._v(" "),s("v-select",{attrs:{items:e.jobCategories,label:"Job Category","single-line":"","item-text":"name","item-value":"val",hint:"What type of work are you interested in?"},model:{value:e.user.jobCategory,callback:function(t){e.$set(e.user,"jobCategory",t)},expression:"user.jobCategory"}}),e._v(" "),s("v-text-field",{attrs:{label:"Job Description",rules:e.jobDescriptionRules,hint:"What type of work are you looking for? Full-time, part-time, contract, big projects, small projects, etc"},model:{value:e.user.jobDescription,callback:function(t){e.$set(e.user,"jobDescription",t)},expression:"user.jobDescription"}}),e._v(" "),s("v-text-field",{attrs:{label:"Skills",rules:e.skillsRules,hint:"What skills do you bring to the table? Use spaces to separate skills."},model:{value:e.user.skills,callback:function(t){e.$set(e.user,"skills",t)},expression:"user.skills"}}),e._v(" "),s("span",{staticClass:"subheading"},[e._v("Coins Accepted")]),s("br"),e._v(" "),s("v-layout",{attrs:{row:"",wrap:""}},e._l(e.user.coins,function(t,n){return s("v-flex",{key:n,staticClass:"xs4 sm3"},[s("v-checkbox",{attrs:{label:e.coinLabels[n]},model:{value:e.user.coins[n],callback:function(t){e.$set(e.user.coins,n,t)},expression:"user.coins[key]"}})],1)})),e._v(" "),s("v-text-field",{directives:[{name:"show",rawName:"v-show",value:e.user.coins.other,expression:"user.coins.other"}],attrs:{label:"Other Coin Name",rules:e.otherCoinRules,hint:"What is the name of the other crypto currency you accept?"},model:{value:e.user.otherCoin,callback:function(t){e.$set(e.user,"otherCoin",t)},expression:"user.otherCoin"}}),e._v(" "),s("v-btn",{attrs:{disabled:!e.valid},on:{click:e.submit}},[e._v("\n          Submit\n        ")]),e._v(" "),s("v-btn",{on:{click:e.clear}},[e._v("Clear")])],1),e._v(" "),s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.sending,expression:"sending"}],attrs:{indeterminate:!0}}),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.errors.length>0,expression:"errors.length > 0"}],attrs:{color:"error",icon:"warning",value:"true"}},e._l(e.errors,function(t,n){return s("p",{key:n},[e._v("\n          "+e._s(t.message)+"\n        ")])}))],1)],1)],1)},we=[],xe={render:ke,staticRenderFns:we},ye=xe,Ce=s("VU/8"),$e=Ce(_e,ye,!1,null,null,null),Le=$e.exports,Re=s("dWJP"),De={data:function(){return{sending:!1,username:"",usernameRules:[function(e){return!!e||"Username is required"}],password:"",passwordRules:[function(e){return!!e||"Password is required"}],errors:[],valid:!0}},methods:{submit:function(){var e=this;this.$refs.form.validate()&&(this.sending=!0,ee.post("/jwt/signin",{},{auth:{username:this.username,password:this.password}}).then(function(t){e.sending=!1;var s=t.headers.authorization.toString().replace("Bearer ","");localStorage.setItem("token",s),ee.defaults.headers.Authorization="Bearer "+s,e.$store.dispatch("setIsLoggedIn",!0),e.$store.dispatch("loadProfile",e.username),e.$router.push({name:"Hello"})}).catch(function(t){e.sending=!1,e.errors.push(t)}))}}},je=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm8:"","offset-sm2":""}},[s("p",{staticClass:"headline"},[e._v("Sign In")]),e._v(" "),s("v-form",{ref:"form",attrs:{"lazy-validation":""},model:{value:e.valid,callback:function(t){e.valid=t},expression:"valid"}},[s("v-text-field",{attrs:{label:"Username",autofocus:"",rules:e.usernameRules,required:""},model:{value:e.username,callback:function(t){e.username=t},expression:"username"}}),e._v(" "),s("v-text-field",{attrs:{label:"Password",rules:e.passwordRules,required:"",type:"password"},on:{keyup:function(t){if(!("button"in t)&&e._k(t.keyCode,"enter",13,t.key))return null;e.submit(t)}},model:{value:e.password,callback:function(t){e.password=t},expression:"password"}}),e._v(" "),s("v-btn",{attrs:{disabled:!e.valid},on:{click:e.submit}},[e._v("\n          Submit\n        ")])],1),e._v(" "),s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.sending,expression:"sending"}],attrs:{indeterminate:!0}}),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.errors.length>0,expression:"errors.length > 0"}],attrs:{color:"error",icon:"warning",value:"true"}},e._l(e.errors,function(t,n){return s("p",{key:n},[e._v("\n          "+e._s(t.message)+"\n        ")])}))],1)],1)],1)},Ie=[],Ve={render:je,staticRenderFns:Ie},Ue=Ve,Pe=s("VU/8"),Se=a,Ee=Pe(De,Ue,!1,Se,null,null),Ae=Ee.exports,Ne={},Fe=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",[s("span",{staticClass:"display-4 primary--text"},[e._v("404'd!")]),e._v(" "),s("p",[e._v("Sorry, we can't find that page.")]),e._v(" "),s("v-btn",{attrs:{large:"",color:"primary",to:"/"}},[s("v-icon",{attrs:{left:""}},[e._v("home")]),e._v("\n      Return Home")],1)],1)],1)],1)},Te=[],He={render:Fe,staticRenderFns:Te},We=He,qe=s("VU/8"),Be=qe(Ne,We,!1,null,null,null),Je=Be.exports,Me=s("NYxO"),Oe={isLoggedIn:!1,user:{},coinLabels:{bcc:"Bitcoin Cash",btc:"Bitcoin",eth:"Ethereum",ltc:"Litecoin",neo:"Neo",other:"Other",xlm:"Lumen",xrp:"Ripple"},jobCategories:[{name:"Hardware",val:"hardware"},{name:"Software",val:"software"},{name:"Writing",val:"writing"},{name:"Legal",val:"legal"},{name:"Labor",val:"labor"},{name:"Automotive",val:"automotive"},{name:"Services",val:"services"},{name:"Others",val:"others"}]},ze={logout:function(e){var t=e.commit;localStorage.removeItem("token"),t("setIsLoggedIn",!1),t("setUser",{})},setIsLoggedIn:function(e,t){(0,e.commit)("setIsLoggedIn",t)},loadProfile:function(e,t){var s=e.commit;ee.get("user/"+t).then(function(e){s("setUser",e.data)}).catch(function(e){console.log(e)})}},Ye={setIsLoggedIn:function(e,t){e.isLoggedIn=t},setUser:function(e,t){e.user=t}},Qe={isLoggedIn:function(e){return e.isLoggedIn},user:function(e){return e.user},jobCategories:function(e){return e.jobCategories},coinLabels:function(e){return e.coinLabels}};i.a.use(Me.a);var Xe=new Me.a.Store({state:Oe,actions:ze,mutations:Ye,getters:Qe});i.a.use(N.a);var Ze=new N.a({mode:"hash",routes:[{path:"/",name:"Hello",component:O},{path:"/Search",name:"Search Listings",meta:{Auth:!0},component:le},{path:"/signup",name:"Signup",component:Le},{path:"/user/:username",name:"User Detail",component:ge},{path:"/about",name:"About",component:Re.default},{path:"/signin",name:"Signin",component:Ae},{path:"/404",component:Je},{path:"*",redirect:"/404"}]});Ze.beforeEach(function(e,t,s){e.matched.some(function(e){return e.meta.Auth})?Xe.state.isLoggedIn?s():s({path:"/signin"}):s()});var Ge=Ze;i.a.config.productionTip=!1,i.a.use(o.a,{components:{VApp:l.a,VNavigationDrawer:u.a,VList:c.f,VBtn:v.a,VIcon:d.a,VCard:m.a,VCheckbox:h.a,VToolbar:f.a,VFooter:p.a,VDivider:b.a,VForm:g.a,VProgressCircular:_.a,VProgressLinear:k.a,VSelect:w.a,VSubheader:x.a,VTextField:y.a,VAlert:C.a,VGrid:$.a,VDialog:L.a,transitions:R.d},directives:{Touch:Touch},theme:{primary:"#ffdb3b",secondary:"#ffab0b",accent:"#757575",error:"#FF5252",info:"#73ea7b",success:"#4CAF50",warning:"#FFC107"}}),new i.a({el:"#app",router:Ge,store:Xe,template:"<App/>",components:{App:A},created:function(){localStorage.getItem("token")&&this.$store.dispatch("setIsLoggedIn",!0)}})},NOHZ:function(e,t){},P0ba:function(e,t){},V5lI:function(e,t){},"X05+":function(e,t){},XC5Q:function(e,t){},acBN:function(e,t){},dWJP:function(e,t,s){"use strict";function n(e){s("8VSw")}var r=s("n4wh"),a=s.n(r),i=s("2V2b"),o=s("VU/8"),l=n,u=o(a.a,i.a,!1,l,null,null);t.default=u.exports},kVeV:function(e,t){},lThp:function(e,t){},n4wh:function(e,t){},pu2v:function(e,t){},"q/9b":function(e,t){},qRVk:function(e,t){},rzhv:function(e,t){},s6ZO:function(e,t){},sBiC:function(e,t){}},["NHnr"]);