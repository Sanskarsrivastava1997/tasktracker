# tasktracker
Task tracker is a project used to track and manage your tasks. In this task, you will build a simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on. 


```
Task Tracker is a tool for managing daily tasks Tasks.
Usage:
        taskTracker <command>
The commands are:
    add                 Add a new Task
    update              Update any existing task
    delete              delete any task
    mark-in-progress    Mark in-progress status to a task
    mark-done           Mark done status to a task
    list                list all tasks

```

<h3 id="example">Example</h3>
<p>The list of commands and their usage is given below:</p>

<div class="not-prose my-6 w-full max-w-full overflow-hidden rounded-lg border border-gray-200"><div class="flex items-center justify-between gap-2 border-b border-gray-200 bg-gray-50 px-3 py-2"><span class="text-sm text-gray-600">bash</span><div class="flex items-center gap-2"><button class="flex size-6 items-center justify-center gap-2 rounded-md text-gray-400 hover:bg-zinc-200 hover:text-black focus:outline-none"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-copy size-3.5" aria-hidden="true"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"></rect><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"></path></svg></button></div></div><div class="mt-0 text-sm [&amp;_pre]:py-0 [&amp;_pre]:grid [&amp;_code]:py-4 [&amp;_code]:w-full [&amp;_code]:grid [&amp;_code]:overflow-x-auto [&amp;_code]:no-scrollbar [&amp;_code]:bg-transparent [&amp;_.line]:px-3 [&amp;_.line]:w-full [&amp;_.line]:relative [&amp;_.line]:min-h-5"><div><pre class="shiki github-light" style="background-color:#fff;color:#24292e" tabindex="0"><code><span class="line"><span style="color:#6A737D"># Adding a new task</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> add</span><span style="color:#032F62"> "Buy groceries"</span></span>
<span class="line"><span style="color:#6A737D"># Output: Task added successfully (ID: 1)</span></span>
<span class="line"></span>
<span class="line"><span style="color:#6A737D"># Updating and deleting tasks</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> update</span><span style="color:#005CC5"> 1</span><span style="color:#032F62"> "Buy groceries and cook dinner"</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> delete</span><span style="color:#005CC5"> 1</span></span>
<span class="line"></span>
<span class="line"><span style="color:#6A737D"># Marking a task as in progress or done</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> mark-in-progress</span><span style="color:#005CC5"> 1</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> mark-done</span><span style="color:#005CC5"> 1</span></span>
<span class="line"></span>
<span class="line"><span style="color:#6A737D"># Listing all tasks</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> list</span></span>
<span class="line"></span>
<span class="line"><span style="color:#6A737D"># Listing tasks by status</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> list</span><span style="color:#032F62"> done</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> list</span><span style="color:#032F62"> todo</span></span>
<span class="line"><span style="color:#6F42C1">task-cli</span><span style="color:#032F62"> list</span><span style="color:#032F62"> in-progress</span></span>
<span class="line"></span></code></pre></div></div></div>


