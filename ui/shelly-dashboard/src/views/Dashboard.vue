<template>
  <div id="app">
    <nav
      class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow"
    >
      <a class="navbar-brand col-md-3 col-lg-2 mr-0 px-3" href="#"
        >Shelly Core</a
      >
      <button
        class="navbar-toggler position-absolute d-md-none collapsed"
        type="button"
        data-toggle="collapse"
        data-target="#sidebarMenu"
        aria-controls="sidebarMenu"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <input
        class="form-control form-control-dark w-100"
        type="text"
        placeholder="Search"
        aria-label="Search"
      />
      <ul class="navbar-nav px-3">
        <li class="nav-item text-nowrap">
          <a class="nav-link" href="#">Sign out</a>
        </li>
      </ul>
    </nav>
    <div class="container-fluid">
      <div class="row">
        <nav
          id="sidebarMenu"
          class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse"
        >
          <div class="sidebar-sticky pt-3">
            <ul class="nav flex-column">
              <li class="nav-item">
                <a class="nav-link" href="/">Dashboard</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/buildings">Buildings</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/rooms">Rooms</a>
              </li>
            </ul>
          </div>
        </nav>
      </div>
    </div>
    <main role="main" class="col-md-9 ml-sm-auto col-lg-10 px-md-4">
      <div id="main" class="container" style="margin-top: 50px">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
};
</script>
