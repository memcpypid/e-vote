<template>
    <header class="bg-[#800000] text-white py-4 px-6 flex justify-between items-center shadow-md">
        <div class="flex items-center space-x-3">
            <img src="/logo-umm.png" alt="UMM Logo" class="w-10 h-10" />
            <h1 class="text-2xl font-bold">E-Vote</h1>
        </div>
        <nav class="space-x-4">
            <a href="#fitur" class="hover:underline">Fitur</a>
            <a href="#cara" class="hover:underline">Panduan</a>
            <router-link to="/login" class="bg-white text-[#800000] px-4 py-2 rounded hover:bg-gray-100 transition">
                Login
            </router-link>
        </nav>
    </header>
</template>
