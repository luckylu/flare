package state

const PAGE_INLINE_STYLE = "/** 月出惊山鸟，时鸣春涧中。**/ * {margin: 0;padding: 0;box-sizing: border-box;}html, body{height: 100%;}body {--color-background: #2d3436;--color-primary: #effbff;--color-accent: #ffa500;--spacing-ui: 10px;--blur: 20px;background-color: var(--color-background);transition: background-color 0.3s;font-family: Roboto, sans-serif;font-size: 14px;background-image: url('/bg.jpg');background-size: cover;background-position: center;background-repeat: no-repeat;backdrop-filter: blur(var(--blur));}input::placeholder {color: var(--color-primary);}h1,h2,h3,h4,h5,div,p,a,span {color: var(--color-primary);}a {text-decoration: none;}.clearfix::after {content: \"\";clear: both;display: table;}h1 {font-weight: 700;font-size: 4em;display: inline-block;}.app-content-uppercase h2 a,.app-content-uppercase .plugin-container h2 {text-transform: uppercase;}h2 a {font-weight: 900;font-size: 20px;margin-bottom: 16px;}.pageview {position: absolute;left: 0;right: 0;top: 0;bottom: 0;display: block;overflow: auto;}#page-home {display: block;overflow: auto;}.container {padding: 25px 40px;width: 90%;margin: 0 auto;}.pull-left {float: left;}.pull-right {float: right;}.no-select {user-select: none;}#container-apps,#container-bookmakrs {margin: 2rem 0;}.plugin-container h2 {font-weight: 900;font-size: 20px;margin-bottom: 16px;}#page-home .footer-container {text-align: center;bottom: 10px;right: 40px;opacity: 0.6;}#page-home .footer-container:hover {opacity: 1;}#page-home.pageview .container {min-width: 375px;}@media (max-width: 1200px) {#page-home.pageview .container {width: 90%;}}@media (max-width: 768px) {#page-home.pageview .container {width: 90%;padding: 20px;margin: 0 auto;}}@media (max-width: 375px) {#page-home.pageview .container {width: 100%;padding: 20px;margin: 0 auto;}}@media (min-width: 1201px) {#page-home.pageview .container {width: 90%;padding: 50px 250px;}}#container-apps .apps-container {width: 100%;}#container-apps .apps-container .app-container {float: left;width: 25%;}#container-apps .apps-container .app-item {display: block;padding: 2px 4px;margin: 12px 0;height: 40px;overflow: hidden;position: relative;}#container-apps .apps-container .app-item:hover {background-color: rgba(0, 0, 0, 0.2);border-radius: 4px;transition: all .1s;}#container-apps .apps-container .app-icon {width: 35px;height: 35px;position: absolute;left: 5px;}#container-apps .apps-container .app-icon * {width: 90%;color: var(--color-primary);}#container-apps .apps-container .app-icon img{display: block;width: 32px;height: 32px;color: var(--color-primary);}.app-content-uppercase #container-apps .apps-container .app-text {text-transform: uppercase;}#container-apps .apps-container .app-text {margin-left: 35px;}#container-apps .apps-container .app-title {font-size: 1em;font-weight: 500;margin-bottom: -4px;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;}#container-apps .apps-container .app-desc {color: var(--color-accent);font-weight: 400;font-size: 0.8em;opacity: 1;margin-top: 4px;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;}@media (max-width: 767px) {#container-apps .apps-container .app-container {width: 50%;}}.app-content-uppercase #container-bookmakrs .bookmark-group-container h3.bookmark-group-title {text-transform: uppercase;}#container-bookmakrs .bookmark-group-container h3.bookmark-group-title {color: var(--color-accent);margin-bottom: 10px;font-size: 16px;font-weight: 400;}#container-bookmakrs .bookmark-group-container {display: inline-block;width: 23.99%;}#container-bookmakrs .bookmark-group-container h3{white-space: nowrap;overflow: hidden;text-overflow: ellipsis;}#container-bookmakrs .bookmark-group-container .bookmark-list {list-style: none;margin: 5px 0;}#container-bookmakrs .bookmark-group-container .bookmark-list a.bookmark {line-height: 2;transition: all 0.25s;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;display: block;}#container-bookmakrs .bookmark-group-container .bookmark-list a.bookmark:hover {text-decoration: underline;padding-left: 10px;}#container-bookmakrs .bookmark-group-container .bookmark-list a.bookmark img {display: inline-block;width: 20px;height: 20px;color: var(--color-primary);vertical-align: middle;}#container-bookmakrs .bookmark-group-container .bookmark svg,#container-bookmakrs .bookmark-group-container .bookmark span {display: inline;vertical-align: middle;}#container-bookmakrs .bookmark-group-container .bookmark span {margin-left: 4px;}#container-bookmakrs .bookmark-group-container .bookmark svg {width: 24px;}@media (max-width: 767px) {#container-bookmakrs .bookmark-group-container {width: 30%;float: none;}}#page-home #hero-container {position: relative;}#plugin-datetime {}#plugin-weather{position: absolute;right: 20px;top: 30px;}#page-home #hero-container p {font-weight: 300;text-transform: uppercase;height: 30px;}#hero-container .weather-value-container {justify-content: center;font-size: 16px;margin-left: 10px;font-weight: 500;}#hero-container .weather-value-container p:first-child {border-bottom: 1px solid var(--color-primary);padding-bottom: 5px;}#search-container #search {width: 100%;padding: 10px 0;color: var(--color-primary);font-size: 20px;background-color: transparent;border: none;border-bottom: 2px solid var(--color-accent);opacity: 0.5;transition: all 0.2s;outline: none;}#search-container #search-label {text-align: right;width: 100%;display: block;opacity: 0;}#search-container .search:active + label {opacity: 1;transition: all 0.2s;}#search-container .search:focus + label {opacity: 1;transition: all 0.2s;}.toolbar-container {width: 38px;position: fixed;bottom: var(--spacing-ui);left: var(--spacing-ui);display: block;-webkit-justify-content: center;justify-content: center;-webkit-align-items: center;align-items: center;visibility: visible;}#btn-open-settings {width: 100%;height: 100%;display: block;text-align: center;overflow: hidden;position: relative;}#btn-open-settings svg {color: var(--color-primary);width: 90%;margin-top: 2px;}#btn-open-settings span {position: absolute;left: 100%;opacity: 0;}#btn-open-help {width: 100%;height: 100%;display: block;text-align: center;overflow: hidden;position: relative;}#btn-open-help svg {color: var(--color-primary);width: 90%;margin-top: 2px;}#btn-open-help span {position: absolute;left: 100%;opacity: 0;}.toolbar-btn-bg{width: 35px;height: 35px;background-color: var(--color-accent);border-radius: 50%;display: block;-webkit-justify-content: center;justify-content: center;-webkit-align-items: center;align-items: center;transition: all 0.3s;visibility: visible;opacity: 0.25;}.toolbar-btn-bg:hover {opacity: 1;}.toolbar-btn-settings{margin: 0 0 10px 0;}}#page-settings .form-group label,#page-settings .form-group select,#page-settings .form-group input.btn-submit,#page-settings .form-group input.option-input {display: block;}#page-settings .form-group select {margin: 8px 0;width: 100%;border: none;border-radius: 4px;padding: 10px;background-color: var(--color-primary);color: var(--color-background);}#page-settings .form-group input.option-input {margin: 8px 0;width: 100%;border: none;border-radius: 4px;padding: 10px;background-color: var(--color-primary);color: var(--color-background);}#page-settings .form-group .help-text {font-size: 12px;color: var(--color-primary);}#page-settings .form-group input.btn-submit {padding: 8px 15px;border: 1px solid var(--color-accent);background-color: var(--color-background);color: var(--color-primary);border-radius: 4px;}#page-settings .container {padding: 25px 20px;width: 100%;margin: 0;}#page-settings h1 {font-size: 30px;}#page-settings .main-container {margin: 20px 0;position: relative;}#page-settings .main-container .setting-group-container {margin-left: 110px;}#page-settings.pageview .container {min-width: 375px;}@media (min-width: 1201px) {#page-settings.pageview .container {width: 90%;padding: 50px 250px;margin: 0 auto;}}@media (max-width: 1200px) {#page-settings.pageview .container {width: 90%;margin: 0 auto;}}@media (max-width: 768px) {#page-settings.pageview .container {width: 90%;padding: 20px;margin: 0 auto;}}@media (max-width: 375px) {#page-settings.pageview .container {width: 100%;padding: 20px;margin: 0 auto;}}.setting-group-container h2 {margin-bottom: 10px;}.setting-group-container .form-group .help-text {margin-bottom: 12px;text-decoration: underline;font-style: italic;opacity: 0.8;}.setting-group-container .form-group input.btn-submit {margin-top: 20px;margin-bottom: 20px;}#page-settings .form-group label {margin-top: 12px;margin-bottom: 6px;}#page-settings .form-group select,#page-settings .form-group input.option-input {margin-top: 6px;margin-bottom: 6px;}#page-settings .form-group textarea {resize: none;width: 100%;height: 80px;border: none;border-radius: 4px;padding: 10px;background-color: var(--color-primary);color: var(--color-background);margin-top: 6px;margin-bottom: 6px;}#page-settings .form-about p {line-height: 2;}#page-settings .form-about a {margin: 0 4px;border-bottom: 1px dotted #ccc;}#page-settings .form-about hr {margin: 10px 0;}#page-settings .sidebar {position: absolute;left: 0;width: 100px;list-style: none;}#page-settings .sidebar a {padding-left: 7px;border-left: 3px solid transparent;display: block;height: 40px;line-height: 40px;transition: all 0.3s;}#page-settings .sidebar a.active {border-left: 3px solid var(--color-primary);}#page-settings .sidebar:hover a.active {border-left: 3px solid transparent;}#page-settings .sidebar a:hover,#page-settings .sidebar a.active:hover {border-left: 3px solid var(--color-primary);}.theme-groups .theme {display: inline-block;width: 30%;}.theme-groups .theme .theme-container {width: 120px;border: 1px solid silver;position: relative;}.theme-groups .theme .theme-color {width: 33.33%;height: 40px;}.theme-groups .theme .theme-title {text-transform: capitalize;margin: 8px 0;color: var(--color-primary);cursor: default;}.theme-groups .theme .theme-choose {position: absolute;z-index: 1;left: 0;right: 0;top: 0;bottom: 0;width: 100%;cursor: pointer;opacity: 0;}.theme-blackboard .theme-color.color-background {background-color: rgb(26, 26, 26);}.theme-blackboard .theme-color.color-primary {background-color: rgb(255, 253, 234);}.theme-blackboard .theme-color.color-accent {background-color: rgb(92, 92, 92);}@media (max-width: 768px) {.theme-groups .theme {display: inline-block;width: 100px;margin-right: 1.5%;}.theme-groups .theme .theme-container {width: 92px;}.theme-groups .theme .theme-color {width: 30px;height: 30px;}}@media (max-width: 375px) {.theme-groups .theme {display: inline-block;width: 47%;margin-right: 1.5%;}.theme-groups .theme .theme-container {width: 92px;}.theme-groups .theme .theme-color {width: 30px;height: 30px;}}"
