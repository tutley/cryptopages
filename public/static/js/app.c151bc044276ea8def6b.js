webpackJsonp([1],{"/626":function(e,t){},"2V2b":function(e,t,s){"use strict";var a=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("div",[e._v("This is the About Page")])])},n=[],r={render:a,staticRenderFns:n};t.a=r},"3TcT":function(e,t){},"6nfk":function(e,t){},"8VSw":function(e,t){},"8dY9":function(e,t){},"90VL":function(e,t){},"9DHQ":function(e,t){},BCfr:function(e,t){},FbEL:function(e,t){},"HNa/":function(e,t){},Ij1S:function(e,t){},JLzs:function(e,t){},MXyr:function(e,t){},"N+IJ":function(e,t){},NHnr:function(e,t,s){"use strict";function a(e){s("8dY9")}function n(e){s("Ij1S")}function r(e){s("MXyr")}Object.defineProperty(t,"__esModule",{value:!0});var i=(s("s6ZO"),s("7+uW")),o=s("M+UZ"),l=s("IcMm"),u=s("rPQa"),c=s("fYhH"),v=s("7Q1V"),d=s("7M7f"),m=s("f/u0"),h=s("z8aH"),f=s("+9ps"),p=s("TWha"),b=s("XRgG"),g=s("DDBM"),_=s("QLUw"),w=s("pqui"),k=s("+1ch"),x=s("7gKz"),y=s("/yi5"),I=s("BOXn"),V=s("JUsQ"),$=s("MPXs"),C=s("6VHA"),L={name:"app",data:function(){return{drawer:!1}},computed:{userIsAuthenticated:function(){return this.$store.getters.isLoggedIn},menuItems:function(){var e=[{icon:"home",title:"Home",link:"/"},{icon:"info",title:"About",link:"/about"},{icon:"lock_open",title:"Sign in",link:"/signin"},{icon:"add_circle",title:"Sign Up",link:"/signup"}];return this.userIsAuthenticated&&(e=[{icon:"home",title:"Home",link:"/"},{icon:"info",title:"About",link:"/about"},{icon:"search",title:"Search Listings",link:"/search"}]),e}},methods:{jump:function(e){this.$router.push(e)},onLogout:function(){this.$store.dispatch("logout"),this.$router.push("/")}},beforeMount:function(){var e=document.getElementById("pwaloader");e&&e.remove()}},D=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-app",{attrs:{id:"cryptopages"}},[s("v-navigation-drawer",{attrs:{app:"",fixed:"",light:""},model:{value:e.drawer,callback:function(t){e.drawer=t},expression:"drawer"}},[s("v-list",[e._l(e.menuItems,function(t){return s("v-list-tile",{key:t.title,attrs:{exact:"",to:t.link}},[s("v-list-tile-action",[s("v-icon",[e._v(e._s(t.icon))])],1),e._v(" "),s("v-list-tile-content",[e._v(e._s(t.title))])],1)}),e._v(" "),e.userIsAuthenticated?s("v-list-tile",{on:{click:e.onLogout}},[s("v-list-tile-action",[s("v-icon",[e._v("exit_to_app")])],1),e._v(" "),s("v-list-tile-content",[e._v("Logout")])],1):e._e()],2)],1),e._v(" "),s("v-toolbar",{attrs:{app:"",dark:"",color:"primary"}},[s("v-toolbar-side-icon",{on:{click:function(t){t.stopPropagation(),e.drawer=!e.drawer}}}),e._v(" "),s("v-toolbar-title",[e._v("\n        Cryptopages\n      ")]),e._v(" "),s("v-spacer"),e._v(" "),s("v-btn",{attrs:{flat:"",to:"/about"}},[e._v("About")])],1),e._v(" "),s("v-content",{ref:"content"},[s("v-fab-transition",[s("v-btn",{attrs:{fab:"",color:"accent",bottom:"",right:"",fixed:"",small:""},on:{click:function(t){e.jump("/search")}}},[s("v-icon",{attrs:{dark:""}},[e._v("search")])],1)],1),e._v(" "),s("div",[e._v(" ")]),e._v(" "),s("router-view"),e._v(" "),s("v-footer",{staticClass:"pa-3"},[s("div",[s("a",{attrs:{href:"https://tomutley.com/"}},[e._v("Tom Utley")])]),e._v(" "),s("v-spacer")],1)],1)],1)},R=[],S={render:D,staticRenderFns:R},U=S,A=s("VU/8"),E=a,N=A(L,U,!1,E,null,null),P=N.exports,j=s("/ocq"),T={name:"hello",data:function(){return{msg:"Its like the Yellow Pages but for Crypto!"}}},F=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",[s("v-flex",[s("div",{staticClass:"hello"},[s("p",{staticClass:"subheading"},[e._v("\n            Hello, welcome to cryptopages.\n            ")]),e._v(" "),s("p",[e._v("\n              This site is not ready to use yet. When it's ready, we'll update this page.\n            ")])])])],1)],1)},q=[],H={render:F,staticRenderFns:q},B=H,M=s("VU/8"),W=n,z=M(T,B,!1,W,null,null),O=z.exports,J=s("Dd8w"),X=s.n(J),Y=s("mtWM"),Q=s.n(Y),Z=null!=localStorage.getItem("token"),G="";Z&&(G="Bearer "+localStorage.getItem("token"));var K={baseURL:"https://cryptopages.club/api/",headers:X()({},Z?{Authorization:G}:{})},ee=Q.a.create(K),te={methods:{displayDetails:function(e){this.$router.push({name:"User Detail",params:{username:e}})}},data:function(){return{userss:[],errors:[],loading:!1}},created:function(){var e=this;this.loading=!0,ee.get("user").then(function(t){e.users=t.data,e.loading=!1}).catch(function(t){e.errors.push(t),e.loading=!1})}},se=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm6:"","offset-sm3":""}},[s("h4",[e._v("Need to refactor this page to be a search page and to show a list of search results")])])],1)],1)},ae=[],ne={render:se,staticRenderFns:ae},re=ne,ie=s("VU/8"),oe=ie(te,re,!1,null,null,null),le=oe.exports,ue=s("woOf"),ce=s.n(ue),ve={data:function(){return{user:{},titleRules:[function(e){return!!e||"Title is required"},function(e){return e&&e.length<=60||"Title mus tbe less than 60 characters"}],bodyRules:[function(e){return!!e||"Body is required"}],errors:[],editable:!1,showDelete:!1,_originaluser:{},loading:!1,processing:!1}},methods:{doEdit:function(e){this._originaluser=ce()({},e),this.editable=!0},cancelEdit:function(){ce()(this.user,this._originaluser),this.editable=!1},saveUpdate:function(){var e=this;this.$refs.form.validate()&&(this.processing=!0,ee.put("user/"+this.$route.params.username,this.user).then(function(t){e.editable=!1,e.processing=!1}).catch(function(t){e.errors.push(t),e.processing=!1}))},doDelete:function(){this.showDelete=!0},cancelDelete:function(){this.showDelete=!1},sendDelete:function(){var e=this;this.processing=!0,ee.delete("user/"+this.$route.params.username).then(function(t){e.processing=!1,e.$router.push({name:"Home"})}).catch(function(t){e.errors.push(t),e.processing=!1})}},created:function(){var e=this;this.loading=!0,ee.get("user/"+this.$route.params.username).then(function(t){e.user=t.data,e.loading=!1}).catch(function(t){e.errors.push(t),e.loading=!1})}},de=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm6:"","offset-sm3":""}},[s("v-card",[s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.loading,expression:"loading"}],attrs:{indeterminate:!0}}),e._v(" "),s("v-form",{ref:"form",attrs:{"lazy-validation":""},model:{value:e.valid,callback:function(t){e.valid=t},expression:"valid"}},[s("v-card-title",{attrs:{headline:""}},[s("h2",{directives:[{name:"show",rawName:"v-show",value:!e.editable,expression:"!editable"}]},[e._v("Need to refactor this Page! "+e._s(e.user.username))]),e._v(" "),s("v-text-field",{directives:[{name:"show",rawName:"v-show",value:e.editable,expression:"editable"}],attrs:{label:"Title",rules:e.titleRules,counter:60,required:""},model:{value:e.user.title,callback:function(t){e.$set(e.user,"title",t)},expression:"user.title"}})],1),e._v(" "),s("v-card-text",[s("p",{directives:[{name:"show",rawName:"v-show",value:!e.editable,expression:"!editable"}]},[e._v(e._s(e.user.body))]),e._v(" "),s("v-text-field",{directives:[{name:"show",rawName:"v-show",value:e.editable,expression:"editable"}],attrs:{label:"Body","multi-line":"",rules:e.bodyRules,required:""},model:{value:e.user.body,callback:function(t){e.$set(e.user,"body",t)},expression:"user.body"}})],1),e._v(" "),s("v-card-actions",{directives:[{name:"show",rawName:"v-show",value:!e.editable,expression:"!editable"}]},[s("v-btn",{attrs:{small:"",color:"accent"},on:{click:function(t){e.doEdit(e.user)}}},[e._v("EDIT")]),e._v(" "),s("v-spacer"),e._v(" "),s("v-btn",{attrs:{small:"",color:"error"},on:{click:e.doDelete}},[e._v("DELETE")])],1),e._v(" "),s("v-card-actions",{directives:[{name:"show",rawName:"v-show",value:e.editable,expression:"editable"}]},[s("v-btn",{attrs:{small:"",color:"success"},on:{click:e.saveUpdate}},[e._v("SAVE")]),e._v(" "),s("v-spacer"),e._v(" "),s("v-btn",{attrs:{small:"",color:"warning"},on:{click:e.cancelEdit}},[e._v("CANCEL")])],1),e._v(" "),s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.processing,expression:"processing"}],attrs:{indeterminate:!0}})],1)],1),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.errors.length>0,expression:"errors.length > 0"}],attrs:{color:"error",icon:"warning",value:"true"}},e._l(e.errors,function(t,a){return s("p",{key:a},[e._v("\n          "+e._s(t.message)+"\n        ")])}))],1)],1),e._v(" "),s("v-dialog",{model:{value:e.showDelete,callback:function(t){e.showDelete=t},expression:"showDelete"}},[s("v-card",[s("v-card-title",{staticClass:"headline"},[e._v("Really delete user?")]),e._v(" "),s("v-card-text",[e._v("Are you sure? Think of the children!")]),e._v(" "),s("v-card-actions",[s("v-spacer"),e._v(" "),s("v-btn",{attrs:{color:"primary",flat:""},on:{click:e.cancelDelete}},[e._v("CANCEL")]),e._v(" "),s("v-btn",{attrs:{color:"error",flat:""},on:{click:e.sendDelete}},[e._v("DELETE")])],1)],1)],1)],1)},me=[],he={render:de,staticRenderFns:me},fe=he,pe=s("VU/8"),be=pe(ve,fe,!1,null,null,null),ge=be.exports,_e=this,we={data:function(){return{sending:!1,usernameAvailable:!0,user:{username:"",password:"",available:!0,name:"",email:{value:"",makePublic:!0},location:{value:"",makePublic:!0},skills:"",jobCategory:"",jobDescription:"",coins:{bcc:!1,btc:!1,eth:!1,ltc:!1,neo:!1,other:!1,xlm:!1,xrp:!1},otherCoin:""},coinLabels:{bcc:"Bitcoin Cash",btc:"Bitcoin",eth:"Ethereum",ltc:"Litecoin",neo:"Neo",other:"Other",xlm:"Lumen",xrp:"Ripple"},jobCategories:[{name:"Hardware",val:"hardware"},{name:"Software",val:"software"},{name:"Writing",val:"writing"},{name:"Legal",val:"legal"},{name:"Labor",val:"labor"},{name:"Automotive",val:"automotive"},{name:"Services",val:"services"},{name:"Others",val:"others"}],usernameRules:[function(e){return!!e||"Username is required"},function(e){return e&&e.length<=60||"Username must be less than 60 characters"},function(e){return _e.usernameAvailable||"Username is not available, sorry"}],passwordRules:[function(e){return!!e||"Password is required"}],nameRules:[],emailRules:[],skillsRules:[],jobDescriptionRules:[],errors:[],valid:!0}},methods:{checkUsername:function(e){var t=this;ee.get("user/checkUsername/"+this.user.username).then(function(e){t.usernameAvailable=!1}).catch(function(e){t.usernameAvailable=!0})},submit:function(){var e=this,t=this.user;t.skills=this.user.skills.split(" "),this.$refs.form.validate()&&(this.sending=!0,ee.post("/user",t).then(function(t){e.sending=!1,console.log(t.data),e.$router.push({name:"User Detail",params:{id:t.data.username}})}).catch(function(t){e.sending=!1,e.errors.push(t)}))},clear:function(){this.$refs.form.reset()}}},ke=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm8:"","offset-sm2":""}},[s("p",{staticClass:"display-2"},[e._v("Sign Up!")]),e._v(" "),s("p",{staticClass:"body-1"},[e._v('Provide as much or as little information as you wish. You must create \n        an account to be able to search and view listings, but you don\'t have to make yourself \n        visible to others. Simply check "available" if you want others to be able to find you!')]),e._v(" "),s("v-form",{ref:"form",attrs:{"lazy-validation":""},model:{value:e.valid,callback:function(t){e.valid=t},expression:"valid"}},[s("v-text-field",{attrs:{label:"Username",rules:e.usernameRules,counter:60,hint:"This is your identifier on this site. It must be unique, and others will see it.",required:""},on:{blur:e.checkUsername},model:{value:e.user.username,callback:function(t){e.$set(e.user,"username",t)},expression:"user.username"}}),e._v(" "),s("v-text-field",{attrs:{label:"Password",rules:e.passwordRules,required:"",type:"password"},model:{value:e.user.password,callback:function(t){e.$set(e.user,"password",t)},expression:"user.password"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Available? (included in search results)"},model:{value:e.user.available,callback:function(t){e.$set(e.user,"available",t)},expression:"user.available"}}),e._v(" "),s("v-text-field",{attrs:{label:"Name",rules:e.nameRules,hint:"Put your name here."},model:{value:e.user.name.value,callback:function(t){e.$set(e.user.name,"value",t)},expression:"user.name.value"}}),e._v(" "),s("v-text-field",{attrs:{label:"Email",rules:e.emailRules,hint:"Let others know the best email to reach you with. You have the option to hide this."},model:{value:e.user.email.value,callback:function(t){e.$set(e.user.email,"value",t)},expression:"user.email.value"}}),e._v(" "),s("v-checkbox",{attrs:{label:"Display Email?"},model:{value:e.user.email.makePublic,callback:function(t){e.$set(e.user.email,"makePublic",t)},expression:"user.email.makePublic"}}),e._v(" "),s("v-select",{attrs:{items:e.jobCategories,label:"Job Category","single-line":"","item-text":"name","item-value":"val","return-object":"",hint:"What type of work are you interested in?"},model:{value:e.user.jobCategory,callback:function(t){e.$set(e.user,"jobCategory",t)},expression:"user.jobCategory"}}),e._v(" "),s("v-text-field",{attrs:{label:"Job Description",rules:e.jobDescriptionRules,hint:"What type of work are you looking for? Full-time, part-time, contract, big projects, small projects, etc"},model:{value:e.user.jobDescription,callback:function(t){e.$set(e.user,"jobDescription",t)},expression:"user.jobDescription"}}),e._v(" "),s("v-text-field",{attrs:{label:"Skills",rules:e.skillsRules,hint:"What skills do you bring to the table? Use spaces to separate skills."},model:{value:e.user.skills,callback:function(t){e.$set(e.user,"skills",t)},expression:"user.skills"}}),e._v(" "),s("span",{staticClass:"subheading"},[e._v("Coins Accepted")]),s("br"),e._v(" "),s("v-layout",{attrs:{row:"",wrap:""}},e._l(e.user.coins,function(t,a){return s("v-flex",{key:a,staticClass:"xs4 sm3"},[s("v-checkbox",{attrs:{label:e.coinLabels[a]},model:{value:e.user.coins[a],callback:function(t){e.$set(e.user.coins,a,t)},expression:"user.coins[key]"}})],1)})),e._v(" "),s("v-btn",{attrs:{disabled:!e.valid},on:{click:e.submit}},[e._v("\n          Submit\n        ")]),e._v(" "),s("v-btn",{on:{click:e.clear}},[e._v("Clear")])],1),e._v(" "),s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.sending,expression:"sending"}],attrs:{indeterminate:!0}}),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.errors.length>0,expression:"errors.length > 0"}],attrs:{color:"error",icon:"warning",value:"true"}},e._l(e.errors,function(t,a){return s("p",{key:a},[e._v("\n          "+e._s(t.message)+"\n        ")])}))],1)],1)],1)},xe=[],ye={render:ke,staticRenderFns:xe},Ie=ye,Ve=s("VU/8"),$e=Ve(we,Ie,!1,null,null,null),Ce=$e.exports,Le=s("dWJP"),De={data:function(){return{sending:!1,username:"",usernameRules:[function(e){return!!e||"Username is required"}],password:"",passwordRules:[function(e){return!!e||"Password is required"}],errors:[],valid:!0}},methods:{submit:function(){var e=this;this.$refs.form.validate()&&(this.sending=!0,ee.post("/jwt/signin",{},{auth:{username:this.username,password:this.password}}).then(function(t){e.sending=!1,console.log(t);var s=t.headers.authorization.toString().replace("Bearer ","");localStorage.setItem("token",s),e.$store.dispatch("setIsLoggedIn",!0),e.$router.push({name:"Hello"})}).catch(function(t){e.sending=!1,e.errors.push(t)}))}}},Re=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",{attrs:{xs12:"",sm8:"","offset-sm2":""}},[s("p",{staticClass:"headline"},[e._v("Sign In")]),e._v(" "),s("v-form",{ref:"form",attrs:{"lazy-validation":""},model:{value:e.valid,callback:function(t){e.valid=t},expression:"valid"}},[s("v-text-field",{attrs:{label:"Username",rules:e.usernameRules,required:""},model:{value:e.username,callback:function(t){e.username=t},expression:"username"}}),e._v(" "),s("v-text-field",{attrs:{label:"Password",rules:e.passwordRules,required:"",type:"password"},model:{value:e.password,callback:function(t){e.password=t},expression:"password"}}),e._v(" "),s("v-btn",{attrs:{disabled:!e.valid},on:{click:e.submit}},[e._v("\n          Submit\n        ")])],1),e._v(" "),s("v-progress-linear",{directives:[{name:"show",rawName:"v-show",value:e.sending,expression:"sending"}],attrs:{indeterminate:!0}}),e._v(" "),s("v-alert",{directives:[{name:"show",rawName:"v-show",value:e.errors.length>0,expression:"errors.length > 0"}],attrs:{color:"error",icon:"warning",value:"true"}},e._l(e.errors,function(t,a){return s("p",{key:a},[e._v("\n          "+e._s(t.message)+"\n        ")])}))],1)],1)],1)},Se=[],Ue={render:Re,staticRenderFns:Se},Ae=Ue,Ee=s("VU/8"),Ne=r,Pe=Ee(De,Ae,!1,Ne,null,null),je=Pe.exports,Te={},Fe=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("v-container",{attrs:{fluid:""}},[s("v-layout",{attrs:{row:""}},[s("v-flex",[s("span",{staticClass:"display-4 primary--text"},[e._v("404'd!")]),e._v(" "),s("p",[e._v("Sorry, we can't find that page.")]),e._v(" "),s("v-btn",{attrs:{large:"",color:"primary",to:"/"}},[s("v-icon",{attrs:{left:""}},[e._v("home")]),e._v("\n      Return Home")],1)],1)],1)],1)},qe=[],He={render:Fe,staticRenderFns:qe},Be=He,Me=s("VU/8"),We=Me(Te,Be,!1,null,null,null),ze=We.exports,Oe=s("NYxO"),Je={isLoggedIn:!1},Xe={logout:function(e){var t=e.commit;localStorage.removeItem("token"),t("setIsLoggedIn",!1)},setIsLoggedIn:function(e,t){(0,e.commit)("setIsLoggedIn",t)}},Ye={setIsLoggedIn:function(e,t){e.isLoggedIn=t}},Qe={isLoggedIn:function(e){return e.isLoggedIn}};i.a.use(Oe.a);var Ze=new Oe.a.Store({state:Je,actions:Xe,mutations:Ye,getters:Qe});i.a.use(j.a);var Ge=new j.a({mode:"history",routes:[{path:"/",name:"Hello",component:O},{path:"/Search",name:"Search Listings",meta:{Auth:!0},component:le},{path:"/signup",name:"Signup",component:Ce},{path:"/user/:username",name:"User Detail",component:ge},{path:"/about",name:"About",component:Le.default},{path:"/signin",name:"Signin",component:je},{path:"/404",component:ze},{path:"*",redirect:"/404"}]});Ge.beforeEach(function(e,t,s){e.matched.some(function(e){return e.meta.Auth})?Ze.state.isLoggedIn?s():s({path:"/signin"}):s()});var Ke=Ge;i.a.config.productionTip=!1,i.a.use(o.a,{components:{VApp:l.a,VNavigationDrawer:u.a,VList:c.f,VBtn:v.a,VIcon:d.a,VCard:m.a,VCheckbox:h.a,VToolbar:f.a,VFooter:p.a,VDivider:b.a,VForm:g.a,VProgressCircular:_.a,VProgressLinear:w.a,VSelect:k.a,VSubheader:x.a,VTextField:y.a,VAlert:I.a,VGrid:V.a,VDialog:$.a,transitions:C.d},directives:{Touch:Touch},theme:{primary:"#ffdb3b",secondary:"#ffab0b",accent:"#757575",error:"#FF5252",info:"#73ea7b",success:"#4CAF50",warning:"#FFC107"}}),new i.a({el:"#app",router:Ke,store:Ze,template:"<App/>",components:{App:P},created:function(){localStorage.getItem("token")&&this.$store.dispatch("setIsLoggedIn",!0)}})},NOHZ:function(e,t){},P0ba:function(e,t){},V5lI:function(e,t){},"X05+":function(e,t){},XC5Q:function(e,t){},acBN:function(e,t){},dWJP:function(e,t,s){"use strict";function a(e){s("8VSw")}var n=s("n4wh"),r=s.n(n),i=s("2V2b"),o=s("VU/8"),l=a,u=o(r.a,i.a,!1,l,null,null);t.default=u.exports},kVeV:function(e,t){},lThp:function(e,t){},n4wh:function(e,t){},pu2v:function(e,t){},"q/9b":function(e,t){},qRVk:function(e,t){},rzhv:function(e,t){},s6ZO:function(e,t){},sBiC:function(e,t){}},["NHnr"]);