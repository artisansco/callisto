<script>
  import { Toaster } from "svelte-sonner";
  import "../app.css";

  const { children } = $props();
</script>

<Toaster richColors position="top-right" closeButton />
{@render children?.()}
