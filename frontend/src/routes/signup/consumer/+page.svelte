<script lang="ts">

    let name = $state('');
    let last_name = $state('');
    let password = $state('');
    let seePassword = $state(false);
    let email = $state('');
    let cell = $state('');
    let landline = $state('');
    
</script>

<style>
    .form-container {
    max-width: 400px;
    margin: 0 auto;
    padding: 16px;
    border: 1px solid #fdfcfcfd;
    border-radius: 16px;
    background: #d3d0d0;
    }

    h1 {
        font-size: 30px;
        margin-top: 16px;
        margin-bottom: 32px;
        text-align: center;
    }

    .password-container {
        display: grid-template-columns;
        align-items: center;
        margin: 0 auto;
    }
    
    .password-label {
        display: block;
        margin-bottom: 0px;
        font-weight: bold;
        border-radius: 8px;
    }

    .password-input {
        width:80%;
        flex-grow: 1;
        margin-right: 10px;
        padding: 10px;
        border: 1px solid #181717;
        border-radius: 8px;
    }

    .toggle-password {
        width:15%;
        background: none;
        color: #007bff; 
        border: 1px solid #181717;
        border-radius: 8px; 
        padding: 10px; 
        cursor: pointer;
        display: inline-flex;
        justify-content: center;
        flex-shrink:0;
    }

    .toggle-password:hover {
        background-color: #84abc5; 
    }

    .toggle-password:focus {
        outline: none; 
        box-shadow: 0 0 4px #613e5e; 
    }

    label {
        display: block;
        margin-bottom: 10px;
        font-weight: bold;
        border-radius: 8px;

    }

    input {
        width:100%;
        padding: 10px;
        margin-bottom: 16px;
        border: 1px solid #181717;
        border-radius: 8px;
    }

    button {
        width: 100%;
        padding: 12px;
        background-color: #007BFF;
        color: white;
        border: none;
        border-radius: 8px;
        margin-top: 16px;
        cursor: pointer;
    }

    button:hover {
    background-color: #295ecf;
    }    

</style>

<h1>Sign Up As Consumer</h1>

<div class="form-container">
    <label for="name">Name</label>
    <input
        type="text"
        id="name"
        bind:value={name}
        placeholder="Enter your name"
    />

    <label for="last_name">Last_name</label>
    <input 
        type="text"
        id="last_name"
        bind:value={last_name}
        placeholder="Enter your last name"
    />

    <label for="password" class="password-label">Password</label>
    <div class="password-container">
        <input
            type={seePassword ? "text" : "password"}
            id="password"
            bind:value={password}
            placeholder="Enter your password"
            class="password-input"
        />
        <button
            type="button"
            class="toggle-password"
            onclick={() => seePassword =!seePassword}
            aria-label="Toggle password visibility"
            >
                {seePassword ? "Hide" : "show"}
        </button>
    </div>

    <label for="email">Email</label>
    <input
        type="email"
        id="email"
        bind:value={email}
        placeholder="Enter your email"
    />

    <label for="cell">Cell</label>
    <input
        type="tel"
        id="cell"
        bind:value={cell}
        placeholder="Enter your number"
    />

    <label for="landline">Landline</label>
    <input
        type="tel"
        id="landline"
        bind:value={landline}
        placeholder="Enter your landline"
    />

    <button type="submit">Sign Up</button>

</div>